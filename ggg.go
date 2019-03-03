package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/BurntSushi/toml"
	"hawx.me/code/ggg/repos"
	"hawx.me/code/ggg/web/assets"
	"hawx.me/code/ggg/web/filters"
	"hawx.me/code/ggg/web/handlers"
	"hawx.me/code/indieauth"
	"hawx.me/code/indieauth/sessions"
	"hawx.me/code/mux"
	"hawx.me/code/route"
	"hawx.me/code/serve"
)

type Conf struct {
	Title  string
	URL    string
	Me     string
	Secret string
	GitDir string
	DbPath string
}

func main() {
	var (
		settingsPath = flag.String("settings", "./settings.toml", "Path to 'settings.toml'")
		port         = flag.String("port", "8080", "Port to run on")
		socket       = flag.String("socket", "", "")
	)
	flag.Parse()

	var conf *Conf
	if _, err := toml.DecodeFile(*settingsPath, &conf); err != nil {
		log.Fatal("toml: ", err)
	}

	db := repos.Open(conf.DbPath, conf.GitDir)
	defer db.Close()

	auth, err := indieauth.Authentication(conf.URL, conf.URL+"/-/callback")
	if err != nil {
		log.Fatal(err)
	}

	session, err := sessions.New(conf.Me, conf.Secret, auth)
	if err != nil {
		log.Fatal(err)
	}

	list := handlers.List(db, conf.Title, conf.URL)
	repo := handlers.Repo(db, conf.Title, conf.URL, session.Choose)

	route.Handle("/", mux.Method{"GET": session.Choose(list.All, list.Public)})

	route.HandleFunc("/:name/*path", func(w http.ResponseWriter, r *http.Request) {
		vars := route.Vars(r)

		if strings.HasSuffix(vars["name"], ".git") {
			repo.Git.ServeHTTP(w, r)
			return
		}

		repo.Html.ServeHTTP(w, r)
	})

	route.Handle("/:name/edit", session.Shield(handlers.Edit(db, conf.Title)))
	route.Handle("/:name/delete", session.Shield(handlers.Delete(db)))

	route.Handle("/-/create", session.Shield(handlers.Create(db, conf.Title)))

	route.HandleFunc("/-/sign-in", session.SignIn())
	route.HandleFunc("/-/callback", session.Callback())
	route.HandleFunc("/-/sign-out", session.SignOut())

	route.Handle("/assets/styles.css", mux.Method{"GET": assets.Styles})
	route.Handle("/assets/highlight.js", mux.Method{"GET": assets.Highlight})
	route.Handle("/assets/filter.js", mux.Method{"GET": assets.Filter})

	serve.Serve(*port, *socket, filters.Log(route.Default))
}

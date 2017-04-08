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
	"hawx.me/code/mux"
	"hawx.me/code/route"
	"hawx.me/code/serve"
	"hawx.me/code/uberich"
)

type Conf struct {
	Title  string
	URL    string
	Secret string
	GitDir string
	DbPath string

	Uberich struct {
		AppName    string
		AppURL     string
		UberichURL string
		Secret     string
	}
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

	store := uberich.NewStore(conf.Secret)
	uberich := uberich.NewClient(conf.Uberich.AppName, conf.Uberich.AppURL, conf.Uberich.UberichURL, conf.Uberich.Secret, store)

	shield := func(h http.Handler) http.Handler {
		return uberich.Protect(h, http.NotFoundHandler())
	}

	list := handlers.List(db, conf.Title, conf.URL)
	repo := handlers.Repo(db, conf.Title, conf.URL, uberich.Protect)

	route.Handle("/", mux.Method{"GET": uberich.Protect(list.All, list.Public)})

	route.HandleFunc("/:name/*path", func(w http.ResponseWriter, r *http.Request) {
		vars := route.Vars(r)

		if strings.HasSuffix(vars["name"], ".git") {
			repo.Git.ServeHTTP(w, r)
			return
		}

		repo.Html.ServeHTTP(w, r)
	})

	route.Handle("/:name/edit", shield(handlers.Edit(db, conf.Title)))
	route.Handle("/:name/delete", shield(handlers.Delete(db)))

	route.Handle("/-/create", shield(handlers.Create(db, conf.Title)))
	route.Handle("/-/sign-in", uberich.SignIn("/"))
	route.Handle("/-/sign-out", uberich.SignOut("/"))

	route.Handle("/assets/styles.css", mux.Method{"GET": assets.Styles})

	serve.Serve(*port, *socket, filters.Log(route.Default))
}

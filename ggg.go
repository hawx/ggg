package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/hawx/persona"
	"github.com/hawx/serve"
	"github.com/stvp/go-toml-config"
	"hawx.me/code/ggg/repos"
	"hawx.me/code/ggg/web/assets"
	"hawx.me/code/ggg/web/filters"
	"hawx.me/code/ggg/web/handlers"
	"hawx.me/code/mux"
	"hawx.me/code/route"
)

var (
	settingsPath = flag.String("settings", "./settings.toml", "Path to 'settings.toml'")
	port         = flag.String("port", "8080", "Port to run on")
	socket       = flag.String("socket", "", "")

	cookieSecret = config.String("secret", "change-me")
	title        = config.String("title", "ggg")
	gitDir       = config.String("gitDir", "./ggg-data/repos")
	dbPath       = config.String("dbPath", "./ggg-data/db")
	user         = config.String("user", "someone@example.com")
	url          = config.String("url", "http://localhost:8080")
)

func main() {
	flag.Parse()

	if err := config.Parse(*settingsPath); err != nil {
		log.Fatal("toml: ", err)
	}

	db := repos.Open(*dbPath, *gitDir)
	defer db.Close()

	store := persona.NewStore(*cookieSecret)
	persona := persona.New(store, *url, []string{*user})

	list := handlers.List(db, *title, *url)
	repo := handlers.Repo(db, *url, persona.Protect)

	route.Handle("/", mux.Method{"GET": persona.Switch(list.All, list.Public)})

	route.HandleFunc("/:name/*path", func(w http.ResponseWriter, r *http.Request) {
		vars := route.Vars(r)

		if strings.HasSuffix(vars["name"], ".git") {
			repo.Git.ServeHTTP(w, r)
			return
		}

		repo.Html.ServeHTTP(w, r)
	})
	route.Handle("/:name/edit", persona.Protect(handlers.Edit(db)))
	route.Handle("/:name/delete", persona.Protect(handlers.Delete(db)))

	route.Handle("/-/create", persona.Protect(handlers.Create(db)))
	route.Handle("/-/sign-in", mux.Method{"POST": persona.SignIn})
	route.Handle("/-/sign-out", mux.Method{"GET": persona.SignOut})

	route.Handle("/assets/styles.css", mux.Method{"GET": assets.Styles})
	route.Handle("/assets/core.js", mux.Method{"GET": assets.Core})

	serve.Serve(*port, *socket, filters.Log(route.Default))
}

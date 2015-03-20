package main

import (
	"net/http"
	"strings"

	"github.com/hawx/ggg/repos"
	"github.com/hawx/ggg/web/assets"
	"github.com/hawx/ggg/web/filters"
	"github.com/hawx/ggg/web/handlers"

	"github.com/hawx/mux"
	"github.com/hawx/persona"
	"github.com/hawx/serve"
	"github.com/stvp/go-toml-config"

	"flag"
	"log"

	"github.com/hawx/route"
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

	dashRouter := route.New()
	dashRouter.Handle("/-/create", persona.Protect(handlers.Create(db)))
	dashRouter.Handle("/-/sign-in", mux.Method{"POST": persona.SignIn})
	dashRouter.Handle("/-/sign-out", mux.Method{"GET": persona.SignOut})

	assetRouter := route.New()
	assetRouter.Handle("/assets/styles.css", mux.Method{"GET": assets.Styles})
	assetRouter.Handle("/assets/core.js", mux.Method{"GET": assets.Core})

	var (
		repo = handlers.Repo(db, *url, persona.Protect)
		list = handlers.List(db, *title, *url)
		edit = persona.Protect(handlers.Edit(db))
		dele = persona.Protect(handlers.Delete(db))
	)

	route.Handle("/", mux.Method{"GET": persona.Switch(list.All, list.Public)})
	route.Handle("/:name", repo.Html)
	route.HandleFunc("/:name/*path", func(w http.ResponseWriter, r *http.Request) {
		vars := route.Vars(r)
		name := vars["name"]
		path := vars["path"]

		if strings.HasSuffix(name, ".git") {
			repo.Git.ServeHTTP(w, r)
			return
		}

		handler, ok := map[string]http.Handler{
			"edit":   edit,
			"delete": dele,
		}[path[1:]]

		if ok {
			handler.ServeHTTP(w, r)
			return
		}

		handler, ok = map[string]http.Handler{
			"-":      dashRouter,
			"assets": assetRouter,
		}[name]

		if ok {
			handler.ServeHTTP(w, r)
			return
		}

		http.NotFound(w, r)
	})

	serve.Serve(*port, *socket, filters.Log(route.Default))
}

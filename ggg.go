package main

import (
	"github.com/hawx/ggg/repos"
	"github.com/hawx/ggg/web/assets"
	"github.com/hawx/ggg/web/filters"
	"github.com/hawx/ggg/web/handlers"

	"github.com/gorilla/mux"
	"github.com/hawx/persona"
	"github.com/hawx/serve"
	"github.com/stvp/go-toml-config"

	"flag"
	"log"
	"net/http"
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

	r := mux.NewRouter()

	list := handlers.List(db, *title, *url)
	r.Methods("GET").Path("/").Handler(persona.Switch(list.All, list.Public))

	create := handlers.Create(db)
	r.Methods("GET").Path("/create").Handler(persona.Protect(create.Get))
	r.Methods("POST").Path("/create").Handler(persona.Protect(create.Post))

	edit := handlers.Edit(db)
	r.Methods("GET").Path("/edit/{name}").Handler(persona.Protect(edit.Get))
	r.Methods("POST").Path("/edit/{name}").Handler(persona.Protect(edit.Post))

	delete := handlers.Delete(db)
	r.Methods("GET").Path("/delete/{name}").Handler(persona.Protect(delete))

	r.Methods("POST").Path("/sign-in").Handler(persona.SignIn)
	r.Methods("GET").Path("/sign-out").Handler(persona.SignOut)

	r.Methods("GET").Path("/assets/styles.css").Handler(assets.Styles)
	r.Methods("GET").Path("/assets/core.js").Handler(assets.Core)

	repo := handlers.Repo(db, *url, persona.Protect)
	r.Methods("GET").PathPrefix("/{name:.+\\.git}/").Handler(repo.Git)
	r.Methods("GET").PathPrefix("/{name}").Handler(repo.Html)

	http.Handle("/", r)

	serve.Serve(*port, *socket, filters.Log(http.DefaultServeMux))
}

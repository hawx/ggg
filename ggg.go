package main

import (
	"github.com/hawx/ggg/assets"
	"github.com/hawx/ggg/repos"
	"github.com/hawx/ggg/views"
	"github.com/hawx/ggg/web/filters"
	"github.com/hawx/ggg/web/handlers"

	"github.com/gorilla/mux"
	"github.com/hawx/persona"
	"github.com/stvp/go-toml-config"

	"flag"
	"log"
	"net/http"
)

var (
	settingsPath = flag.String("settings", "./settings.toml", "Path to 'settings.toml'")
	port         = flag.String("port", "8080", "Port to run on")

	audience     = config.String("audience", "localhost")
	cookieSecret = config.String("secret", "change-me")
	title        = config.String("title", "ggg")
	gitDir       = config.String("gitDir", "./ggg-data/repos")
	dbPath       = config.String("dbPath", "./ggg-data/db")
	user         = config.String("user", "someone@example.com")
	url          = config.String("url", "http://localhost:8080")
	token        = config.String("token", "")
)

type Ctx struct {
	Title       string
	Url         string
	Repos       []repos.Repo
}

func List(db repos.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repos := []repos.Repo{}
		for _, repo := range db.GetAll() {
			if !repo.IsPrivate {
				repos = append(repos, repo)
			}
		}

		w.Header().Add("Content-Type", "text/html")
		views.List.Execute(w, Ctx{*title, *url, repos})
	})
}

func Admin(db repos.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		views.Admin.Execute(w, Ctx{*title, *url, db.GetAll()})
	})
}


func main() {
	flag.Parse()

	if err := config.Parse(*settingsPath); err != nil {
		log.Fatal("toml: ", err)
	}

	db := repos.Open(*dbPath, *gitDir)
	defer db.Close()

	store := persona.NewStore(*cookieSecret)
	persona := persona.New(store, *audience, []string{*user})

	r := mux.NewRouter()
	r.Methods("GET").Path("/").Handler(persona.Switch(Admin(db), List(db)))

	github := handlers.GitHub(db, *token)
	r.Methods("GET").Path("/github").Handler(github)

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

	repo := handlers.Repo(db, *url)
	r.Methods("GET").PathPrefix("/{name:.+\\.git}/").Handler(repo.Git)
	r.Methods("GET").PathPrefix("/{name}").Handler(repo.Html)

	http.Handle("/", r)

	log.Println("Running on :" + *port)
	log.Fatal(http.ListenAndServe(":"+*port, filters.Log(http.DefaultServeMux)))
}

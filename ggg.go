package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hawx/ggg/assets"
	"github.com/hawx/ggg/repos"
	"github.com/hawx/ggg/views"
	asset "github.com/hawx/wwwhat/assets"
	"github.com/hawx/wwwhat/persona"
	"github.com/stvp/go-toml-config"
	"log"
	"net/http"
)

var (
	settingsPath = flag.String("settings", "./settings.toml", "Path to 'settings.toml'")
	port         = flag.String("port", "8080", "Port to run on")

	audience     = config.String("audience", "localhost")
	cookieSecret = config.String("secret", "change-me")
	title        = config.String("title", "git")
	description  = config.String("description", "My own, personal git-server.")
	gitDir       = config.String("gitDir", "./ggg-git")
	dbPath       = config.String("dbPath", "./ggg-db")
	user         = config.String("user", "someone@example.com")
	url          = config.String("url", "http://localhost:8080")
)

type Ctx struct {
	Title       string
	Description string
	Url         string
	Repos       repos.Repos
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func List(db repos.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repos := repos.Repos{}
		for _, repo := range db.GetAll() {
			if !repo.IsPrivate {
				repos = append(repos, repo)
			}
		}

		body := views.List.Render(Ctx{*title, *description, *url, repos})
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, body)
	})
}

func Admin(db repos.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := views.Admin.Render(Ctx{*title, *description, *url, db.GetAll()})
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, body)
	})
}

func CreateGet(db repos.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := views.Create.Render()
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, body)
	})
}

func CreatePost(db repos.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db.Create(r.FormValue("name"), r.FormValue("web"), r.FormValue("description"), r.FormValue("private") == "private")
		http.Redirect(w, r, "/", 302)
	})
}

func EditGet(db repos.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repoName := mux.Vars(r)["name"]
		repo := db.Get(repoName)

		body := views.Edit.Render(repo)
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, body)
	})
}

func EditPost(db repos.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repoName := mux.Vars(r)["name"]
		repo := db.Get(repoName)

		repo.Description = r.FormValue("description")
		repo.Web = r.FormValue("web")
		repo.IsPrivate = r.FormValue("private") == "private"
		db.Save(repo)
		http.Redirect(w, r, "/", 302)
	})
}

func Delete(db repos.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repoName := mux.Vars(r)["name"]
		db.Delete(db.Get(repoName))
		http.Redirect(w, r, "/", 302)
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
	protect := persona.Protector(strore, []string{*user})
	cond := persona.Conditional(store, []string{*user})

	r := mux.NewRouter()
	r.Methods("GET").Path("/").Handler(cond(Admin(db), List(db)))
	r.Methods("GET").Path("/create").Handler(protect(CreateGet(db)))
	r.Methods("GET").Path("/create").Handler(protect(CreatePost(db)))
	r.Methods("GET").Path("/edit/{name}").Handler(protect(EditGet(db)))
	r.Methods("POST").Path("/edit/{name}").Handler(protect(EditPost(db)))
	r.Methods("GET").Path("/delete/{name}").Handler(protect(Delete(db)))
	r.Methods("POST").Path("/sign-in").Handler(persona.SignIn(store, *audience))
	r.Methods("GET").Path("/sign-out").Handler(persona.SignOut(store))

	http.Handle("/", r)
	http.Handle("/assets/", http.StripPrefix("/assets/", asset.Server(map[string]string{
		"styles.css": assets.Styles,
		"core.js":    assets.Core,
	})))

	log.Println("Running on :" + *port)
	log.Fatal(http.ListenAndServe(":"+*port, Log(http.DefaultServeMux)))
}

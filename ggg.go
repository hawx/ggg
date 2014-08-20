package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/context"
	"github.com/hawx/ggg/assets"
	"github.com/hawx/ggg/repos"
	"github.com/hawx/ggg/views"
	asset "github.com/hawx/wwwhat/assets"
	"github.com/hawx/wwwhat/persona"
	"github.com/stvp/go-toml-config"
	"log"
	"net/http"
	"path/filepath"
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
	user         = config.String("user", "m@hawx.me")
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

func Create(db repos.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			body := views.Create.Render()
			w.Header().Add("Content-Type", "text/html")
			fmt.Fprintf(w, body)
		} else if r.Method == "POST" {
			r.ParseForm()
			db.Create(r.PostForm["name"][0], r.PostForm["web"][0], r.PostForm["description"][0], len(r.PostForm["private"]) != 0)
			http.Redirect(w, r, "/", 302)
		}
	})
}

func Edit(db repos.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repoName := filepath.Base(r.URL.Path)
		repo := db.Get(repoName)

		if r.Method == "GET" {
			body := views.Edit.Render(repo)
			w.Header().Add("Content-Type", "text/html")
			fmt.Fprintf(w, body)
		} else if r.Method == "POST" {
			r.ParseForm()
			repo.Description = r.PostForm["description"][0]
			repo.Web = r.PostForm["web"][0]
			repo.IsPrivate = len(r.PostForm["private"]) != 0
			db.Save(repo)
			http.Redirect(w, r, "/", 302)
		}
	})
}

func Delete(db repos.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			repoName := filepath.Base(r.URL.Path)
			db.Delete(db.Get(repoName))
			http.Redirect(w, r, "/", 302)
		}
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
	protect := persona.Conditional(store, []string{*user})

	http.Handle("/", protect(Admin(db), List(db)))
	http.Handle("/create", protect(Create(db), http.NotFoundHandler()))
	http.Handle("/edit/", protect(Edit(db), http.NotFoundHandler()))
	http.Handle("/delete/", protect(Delete(db), http.NotFoundHandler()))
	http.Handle("/sign-in", persona.SignIn(store, *audience))
	http.Handle("/sign-out", persona.SignOut(store))
	http.Handle("/assets/", http.StripPrefix("/assets/", asset.Server(map[string]string{
		"styles.css": assets.Styles,
		"core.js":    assets.Core,
	})))

	log.Println("Running on :8080")
	log.Fatal(http.ListenAndServe(":8080", context.ClearHandler(Log(http.DefaultServeMux))))
}

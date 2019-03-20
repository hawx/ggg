package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"strings"

	"hawx.me/code/ggg/handlers"
	"hawx.me/code/ggg/repos"
	"hawx.me/code/indieauth"
	"hawx.me/code/indieauth/sessions"
	"hawx.me/code/mux"
	"hawx.me/code/route"
	"hawx.me/code/serve"
)

func main() {
	var (
		title   = flag.String("title", "ggg", "")
		url     = flag.String("url", "http://localhost:8080", "")
		me      = flag.String("me", "", "")
		secret  = flag.String("secret", "plschange", "")
		gitDir  = flag.String("git-dir", "./_git_repos", "")
		dbPath  = flag.String("db", "./db", "")
		port    = flag.String("port", "8080", "Port to run on")
		socket  = flag.String("socket", "", "")
		webPath = flag.String("web", "web", "")
	)
	flag.Parse()

	auth, err := indieauth.Authentication(*url, *url+"/-/callback")
	if err != nil {
		log.Fatal(err)
	}

	session, err := sessions.New(*me, *secret, auth)
	if err != nil {
		log.Fatal(err)
	}

	templates, err := template.ParseGlob(*webPath + "/template/*.gotmpl")
	if err != nil {
		log.Fatal(err)
	}

	db := repos.Open(*dbPath, *gitDir)
	defer db.Close()

	list := handlers.List(db, *title, *url, templates)
	repo := handlers.Repo(db, *title, *url, session.Choose, templates)

	route.Handle("/", mux.Method{"GET": session.Choose(list.All, list.Public)})

	route.HandleFunc("/:name/*path", func(w http.ResponseWriter, r *http.Request) {
		vars := route.Vars(r)

		if strings.HasSuffix(vars["name"], ".git") {
			repo.Git.ServeHTTP(w, r)
			return
		}

		repo.Html.ServeHTTP(w, r)
	})

	route.Handle("/:name/edit", session.Shield(handlers.Edit(db, *title, templates)))
	route.Handle("/:name/delete", session.Shield(handlers.Delete(db)))

	route.Handle("/-/create", session.Shield(handlers.Create(db, *title, templates)))

	route.HandleFunc("/-/sign-in", session.SignIn())
	route.HandleFunc("/-/callback", session.Callback())
	route.HandleFunc("/-/sign-out", session.SignOut())

	route.Handle("/public/*path", http.StripPrefix("/public", http.FileServer(http.Dir(*webPath+"/static"))))

	serve.Serve(*port, *socket, Log(route.Default))
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

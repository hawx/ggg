package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"strings"

	"hawx.me/code/ggg/handlers"
	"hawx.me/code/ggg/repos"
	"hawx.me/code/indieauth/v2"
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

	sessions, err := indieauth.NewSessions(*secret, &indieauth.Config{
		ClientID:    *url,
		RedirectURL: *url + "/-/callback",
	})
	if err != nil {
		log.Fatal(err)
	}

	templates, err := template.ParseGlob(*webPath + "/template/*.gotmpl")
	if err != nil {
		log.Fatal(err)
	}

	db := repos.Open(*dbPath, *gitDir)
	defer db.Close()

	choose := func(signedIn, signedOut http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if _, ok := sessions.SignedIn(r); ok {
				signedIn.ServeHTTP(w, r)
			} else {
				signedOut.ServeHTTP(w, r)
			}
		}
	}

	shield := func(h http.Handler) http.Handler {
		return choose(h, http.NotFoundHandler())
	}

	list := handlers.List(db, *title, *url, templates)
	repo := handlers.Repo(db, *title, *url, choose, templates)

	route.Handle("/", mux.Method{"GET": choose(list.All, list.Public)})

	route.HandleFunc("/:name/*path", func(w http.ResponseWriter, r *http.Request) {
		vars := route.Vars(r)

		if strings.HasSuffix(vars["name"], ".git") {
			repo.Git.ServeHTTP(w, r)
			return
		}

		repo.Html.ServeHTTP(w, r)
	})

	route.Handle("/:name/edit", shield(handlers.Edit(db, *title, templates)))
	route.Handle("/:name/delete", shield(handlers.Delete(db)))

	route.Handle("/-/create", shield(handlers.Create(db, *title, templates)))

	route.HandleFunc("/-/sign-in", func(w http.ResponseWriter, r *http.Request) {
		if err := sessions.RedirectToSignIn(w, r, *me); err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
		}
	})

	route.HandleFunc("/-/callback", func(w http.ResponseWriter, r *http.Request) {
		if err := sessions.Verify(w, r); err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
	})

	route.HandleFunc("/-/sign-out", func(w http.ResponseWriter, r *http.Request) {
		if err := sessions.SignOut(w, r); err != nil {
			log.Println(err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
	})

	route.Handle("/public/*path", http.StripPrefix("/public", http.FileServer(http.Dir(*webPath+"/static"))))

	serve.Serve(*port, *socket, route.Default)
}

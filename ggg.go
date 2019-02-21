package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/sessions"
	"hawx.me/code/ggg/repos"
	"hawx.me/code/ggg/web/assets"
	"hawx.me/code/ggg/web/filters"
	"hawx.me/code/ggg/web/handlers"
	"hawx.me/code/indieauth"
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

	endpoints, err := indieauth.FindEndpoints(conf.Me)
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(conf.Secret)

	protect := func(good, bad http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if addr := store.Get(r); addr == conf.Me {
				good.ServeHTTP(w, r)
			} else {
				bad.ServeHTTP(w, r)
			}
		})
	}

	shield := func(h http.Handler) http.Handler {
		return protect(h, http.NotFoundHandler())
	}

	list := handlers.List(db, conf.Title, conf.URL)
	repo := handlers.Repo(db, conf.Title, conf.URL, protect)

	route.Handle("/", mux.Method{"GET": protect(list.All, list.Public)})

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

	route.HandleFunc("/-/sign-in", func(w http.ResponseWriter, r *http.Request) {
		state, err := store.SetState(w, r)
		if err != nil {
			http.Error(w, "could not start auth", http.StatusInternalServerError)
			return
		}

		redirectURL := auth.RedirectURL(endpoints, conf.Me, state)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, redirectURL, http.StatusFound)
	})

	route.HandleFunc("/-/callback", func(w http.ResponseWriter, r *http.Request) {
		state := store.GetState(r)

		if r.FormValue("state") != state {
			http.Error(w, "state is bad", http.StatusBadRequest)
			return
		}

		me, err := auth.Exchange(endpoints, r.FormValue("code"))
		if err != nil || me != conf.Me {
			http.Error(w, "nope", http.StatusForbidden)
			return
		}

		store.Set(w, r, me)
		http.Redirect(w, r, "/", http.StatusFound)
	})

	route.HandleFunc("/-/sign-out", func(w http.ResponseWriter, r *http.Request) {
		store.Set(w, r, "")
		http.Redirect(w, r, "/", http.StatusFound)
	})

	route.Handle("/assets/styles.css", mux.Method{"GET": assets.Styles})
	route.Handle("/assets/highlight.js", mux.Method{"GET": assets.Highlight})
	route.Handle("/assets/filter.js", mux.Method{"GET": assets.Filter})

	serve.Serve(*port, *socket, filters.Log(route.Default))
}

type meStore struct {
	store sessions.Store
}

func NewStore(secret string) *meStore {
	return &meStore{sessions.NewCookieStore([]byte(secret))}
}

func (s meStore) Get(r *http.Request) string {
	session, _ := s.store.Get(r, "session")

	if v, ok := session.Values["me"].(string); ok {
		return v
	}

	return ""
}

func (s meStore) Set(w http.ResponseWriter, r *http.Request, me string) {
	session, _ := s.store.Get(r, "session")
	session.Values["me"] = me
	session.Save(r, w)
}

func (s meStore) SetState(w http.ResponseWriter, r *http.Request) (string, error) {
	bytes := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(bytes)

	session, _ := s.store.Get(r, "session")
	session.Values["state"] = state
	return state, session.Save(r, w)
}

func (s meStore) GetState(r *http.Request) string {
	session, _ := s.store.Get(r, "session")

	if v, ok := session.Values["state"].(string); ok {
		return v
	}

	return ""
}

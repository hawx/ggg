package handlers

import (
	"github.com/hawx/ggg/repos"
	"github.com/hawx/ggg/web/views"

	"github.com/gorilla/mux"
	"github.com/hawx/persona"

	"net/http"
)

func Repo(db repos.Db, url string, protect persona.Filter) RepoHandler {
	h := repoHandler{db}

	return RepoHandler{
		Html: h.Html(url, protect),
		Git:  h.Git(protect),
	}
}

type RepoHandler struct {
	Html http.Handler
	Git  http.Handler
}

type repoHandler struct {
	db repos.Db
}

func (h repoHandler) Git(protect persona.Filter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		repoName := name[:len(name)-4]
		repo := h.db.Get(repoName)

		if repo.IsPrivate {
			http.NotFound(w, r)
			return
		}

		http.StripPrefix("/"+name+"/",
			http.FileServer(http.Dir(repo.Path))).ServeHTTP(w, r)
	})
}

func (h repoHandler) Html(url string, protect persona.Filter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := mux.Vars(r)["name"]
		repo := h.db.Get(name)

		if repo.Name == "" {
			http.NotFound(w, r)
			return
		}

		innerHandler := h.htmlPage(repo, url)

		if repo.IsPrivate {
			protect(innerHandler).ServeHTTP(w, r)
			return
		}

		innerHandler.ServeHTTP(w, r)
	})
}

func (h repoHandler) htmlPage(repo *repos.Repo, url string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		views.Repo.Execute(w, struct {
			*repos.Repo
			Url string
		}{repo, url})
	})
}

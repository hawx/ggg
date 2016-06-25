package handlers

import (
	"hawx.me/code/ggg/repos"
	"hawx.me/code/ggg/web/views"

	"hawx.me/code/route"

	"net/http"
)

type Middleware func(http.Handler) http.Handler

func Repo(db repos.Db, url string, shield Middleware) RepoHandler {
	h := repoHandler{db}

	return RepoHandler{
		Html: h.Html(url, shield),
		Git:  h.Git(shield),
	}
}

type RepoHandler struct {
	Html http.Handler
	Git  http.Handler
}

type repoHandler struct {
	db repos.Db
}

func (h repoHandler) Git(shield Middleware) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := route.Vars(r)["name"]
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

func (h repoHandler) Html(url string, shield Middleware) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := route.Vars(r)["name"]
		repo := h.db.Get(name)

		if repo.Name == "" {
			http.NotFound(w, r)
			return
		}

		innerHandler := h.htmlPage(repo, url)

		if repo.IsPrivate {
			shield(innerHandler).ServeHTTP(w, r)
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

package handlers

import (
	"hawx.me/code/ggg/repos"
	"hawx.me/code/ggg/web/views"

	"hawx.me/code/route"

	"net/http"
)

type Protect func(handler, errHandler http.Handler) http.Handler

func Repo(db repos.Db, title, url string, protect Protect) RepoHandler {
	h := repoHandler{db}

	return RepoHandler{
		Html: h.Html(title, url, protect),
		Git:  h.Git(),
	}
}

type RepoHandler struct {
	Html http.Handler
	Git  http.Handler
}

type repoHandler struct {
	db repos.Db
}

func (h repoHandler) Git() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := route.Vars(r)["name"]
		repoName := name[:len(name)-4]
		repo := h.db.Get(repoName)

		if repo.Name == "" || repo.IsPrivate {
			http.NotFound(w, r)
			return
		}

		http.StripPrefix("/"+name+"/",
			http.FileServer(http.Dir(repo.Path))).ServeHTTP(w, r)
	})
}

func (h repoHandler) Html(title, url string, protect Protect) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := route.Vars(r)["name"]
		repo := h.db.Get(name)

		if repo.Name == "" {
			http.NotFound(w, r)
			return
		}

		innerHandler := h.htmlPage(title, repo, url)

		if repo.IsPrivate {
			protect(innerHandler(true), http.NotFoundHandler()).ServeHTTP(w, r)
			return
		}

		protect(innerHandler(true), innerHandler(false)).ServeHTTP(w, r)
	})
}

func (h repoHandler) htmlPage(title string, repo *repos.Repo, url string) func(bool) http.Handler {
	return func(loggedIn bool) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			views.Repo.Execute(w, struct {
				Title string
				*repos.Repo
				Url      string
				LoggedIn bool
			}{title, repo, url, loggedIn})
		})
	}
}

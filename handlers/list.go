package handlers

import (
	"net/http"

	"hawx.me/code/ggg/repos"
)

func List(db repos.Db, title, url string, templates Templates) ListHandler {
	h := listHandler{db, title, url, templates}

	return ListHandler{
		Public: h.Public(),
		All:    h.All(),
	}
}

type ListHandler struct {
	Public http.Handler
	All    http.Handler
}

type listHandler struct {
	db        repos.Db
	title     string
	url       string
	templates Templates
}

type Ctx struct {
	Title    string
	Url      string
	Repos    []*repos.Repo
	LoggedIn bool
}

func (h listHandler) Public() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repos := []*repos.Repo{}
		for _, repo := range h.db.GetAll() {
			if !repo.IsPrivate {
				repos = append(repos, repo)
			}
		}

		w.Header().Add("Content-Type", "text/html")
		h.templates.ExecuteTemplate(w, "list.gotmpl", Ctx{h.title, h.url, repos, false})
	})
}

func (h listHandler) All() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		h.templates.ExecuteTemplate(w, "list.gotmpl", Ctx{h.title, h.url, h.db.GetAll(), true})
	})
}

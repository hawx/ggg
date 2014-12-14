package handlers

import (
	"github.com/hawx/ggg/repos"
	"github.com/hawx/ggg/web/views"

	"net/http"
)

func List(db repos.Db, title, url string) ListHandler {
	h := listHandler{db, title, url}

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
	db    repos.Db
	title string
	url   string
}

type Ctx struct {
	Title string
	Url   string
	Repos []*repos.Repo
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
		views.List.Execute(w, Ctx{h.title, h.url, repos})
	})
}

func (h listHandler) All() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		views.Admin.Execute(w, Ctx{h.title, h.url, h.db.GetAll()})
	})
}

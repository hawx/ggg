package handlers

import (
	"github.com/hawx/ggg/git"
	"github.com/hawx/ggg/repos"
	"github.com/hawx/ggg/web/views"

	"net/http"
)

func Create(db repos.Db) CreateHandler {
	h := createHandler{db}

	return CreateHandler{
		Get:  h.Get(),
		Post: h.Post(),
	}
}

type CreateHandler struct {
	Get  http.Handler
	Post http.Handler
}

type createHandler struct {
	db repos.Db
}

func (h createHandler) Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		views.Create.Execute(w, nil)
	})
}

func (h createHandler) Post() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.db.Create(
			r.FormValue("name"),
			r.FormValue("web"),
			r.FormValue("description"),
			"",
			r.FormValue("private") == "private")

		created := h.db.Get(r.FormValue("name"))
		git.CreateRepo(created.Path)

		http.Redirect(w, r, "/", 302)
	})
}

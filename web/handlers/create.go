package handlers

import (
	"hawx.me/code/ggg/git"
	"hawx.me/code/ggg/repos"
	"hawx.me/code/ggg/web/views"
	"hawx.me/code/mux"

	"net/http"
)

func Create(db repos.Db) http.Handler {
	h := createHandler{db}

	return mux.Method{
		"GET":  h.Get(),
		"POST": h.Post(),
	}
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

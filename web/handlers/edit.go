package handlers

import (
	"hawx.me/code/ggg/repos"
	"hawx.me/code/ggg/web/views"

	"hawx.me/code/mux"
	"hawx.me/code/route"

	"net/http"
)

func Edit(db repos.Db) http.Handler {
	h := editHandler{db}

	return mux.Method{
		"GET":  h.Get(),
		"POST": h.Post(),
	}
}

type editHandler struct {
	db repos.Db
}

func (h editHandler) Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repoName := route.Vars(r)["name"]
		repo := h.db.Get(repoName)

		w.Header().Add("Content-Type", "text/html")
		views.Edit.Execute(w, repo)
	})
}

func (h editHandler) Post() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repoName := route.Vars(r)["name"]
		repo := h.db.Get(repoName)

		repo.Web = r.FormValue("web")
		repo.Description = r.FormValue("description")
		repo.Branch = r.FormValue("branch")
		repo.IsPrivate = r.FormValue("private") == "private"
		h.db.Save(repo)
		http.Redirect(w, r, "/", 302)
	})
}

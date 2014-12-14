package handlers

import (
	"github.com/hawx/ggg/repos"
	"github.com/hawx/ggg/web/views"

	"github.com/gorilla/mux"

	"net/http"
)

func Edit(db repos.Db) EditHandler {
	h := editHandler{db}

	return EditHandler{
		Get:  h.Get(),
		Post: h.Post(),
	}
}

type EditHandler struct {
	Get  http.Handler
	Post http.Handler
}

type editHandler struct {
	db repos.Db
}

func (h editHandler) Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repoName := mux.Vars(r)["name"]
		repo := h.db.Get(repoName)

		w.Header().Add("Content-Type", "text/html")
		views.Edit.Execute(w, repo)
	})
}

func (h editHandler) Post() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repoName := mux.Vars(r)["name"]
		repo := h.db.Get(repoName)

		repo.Web = r.FormValue("web")
		repo.Description = r.FormValue("description")
		repo.Branch = r.FormValue("branch")
		repo.IsPrivate = r.FormValue("private") == "private"
		h.db.Save(repo)
		http.Redirect(w, r, "/", 302)
	})
}

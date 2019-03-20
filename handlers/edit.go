package handlers

import (
	"net/http"

	"hawx.me/code/ggg/repos"
	"hawx.me/code/mux"
	"hawx.me/code/route"
)

func Edit(db repos.Db, title string, templates Templates) http.Handler {
	h := editHandler{db, title, templates}

	return mux.Method{
		"GET":  h.Get(),
		"POST": h.Post(),
	}
}

type editHandler struct {
	db        repos.Db
	title     string
	templates Templates
}

func (h editHandler) Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repoName := route.Vars(r)["name"]
		repo := h.db.Get(repoName)

		w.Header().Add("Content-Type", "text/html")
		h.templates.ExecuteTemplate(w, "edit.gotmpl", struct {
			Title string
			*repos.Repo
			LoggedIn bool
		}{h.title, repo, true})
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
		http.Redirect(w, r, "/"+repoName, 302)
	})
}

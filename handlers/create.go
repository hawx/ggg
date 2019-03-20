package handlers

import (
	"net/http"

	"hawx.me/code/ggg/repos"
	"hawx.me/code/mux"
)

func Create(db repos.Db, title string, templates Templates) http.Handler {
	h := createHandler{db, title, templates}

	return mux.Method{
		"GET":  h.Get(),
		"POST": h.Post(),
	}
}

type createHandler struct {
	db        repos.Db
	title     string
	templates Templates
}

func (h createHandler) Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		h.templates.ExecuteTemplate(w, "create.gotmpl", struct {
			Title    string
			LoggedIn bool
		}{h.title, true})
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
		repos.CreateRepo(created.Path)

		http.Redirect(w, r, "/", 302)
	})
}

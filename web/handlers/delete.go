package handlers

import (
	"github.com/hawx/ggg/repos"

	"github.com/gorilla/mux"

	"net/http"
)

func Delete(db repos.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repoName := mux.Vars(r)["name"]
		db.Delete(db.Get(repoName))
		http.Redirect(w, r, "/", 302)
	})
}

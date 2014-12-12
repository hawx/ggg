package handlers

import (
	"github.com/hawx/ggg/repos"

	"github.com/gorilla/mux"

	"net/http"
	"os"
)

func Delete(db repos.Db) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repoName := mux.Vars(r)["name"]

		repo := db.Get(repoName)
		db.Delete(repo)
		os.RemoveAll(repo.Path)

		http.Redirect(w, r, "/", 302)
	})
}

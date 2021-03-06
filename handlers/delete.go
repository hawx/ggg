package handlers

import (
	"hawx.me/code/ggg/repos"

	"hawx.me/code/mux"
	"hawx.me/code/route"

	"net/http"
	"os"
)

func Delete(db repos.Db) http.Handler {
	return mux.Method{
		"GET": http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			repoName := route.Vars(r)["name"]

			repo := db.Get(repoName)
			db.Delete(repo)
			os.RemoveAll(repo.Path)

			http.Redirect(w, r, "/", 302)
		}),
	}
}

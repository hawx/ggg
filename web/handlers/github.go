package handlers

import (
	"github.com/hawx/ggg/repos"
	"github.com/hawx/ggg/git"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"

	"net/http"
	"log"
)

func GitHub(db repos.Db, token string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := &oauth.Transport{
		  Token: &oauth.Token{AccessToken: token},
		}

		client := github.NewClient(t.Client())

		// list all repositories for the authenticated user
		repos, _, _ := client.Repositories.List("", &github.RepositoryListOptions{
		  Type: "owner",
		})

		for _, repo := range repos {
			copyRepo(db, repo)
		}

		http.Redirect(w, r, "/", 302)
	})
}

func copyRepo(db repos.Db, repo github.Repository) {
	if db.Get(*repo.Name).Name != "" {
		log.Println("skipping", *repo.Name)
		return
	}

	homepage := ""
	if repo.Homepage != nil {
		homepage = *repo.Homepage
	}

	description := ""
	if repo.Description != nil {
		description = *repo.Description
	}

	private := false
	if repo.Private != nil {
		private = *repo.Private
	}

	db.Create(
		*repo.Name,
		homepage,
		description,
		private)

	created := db.Get(*repo.Name)

	git.CopyRepo(created.Path, *repo.CloneURL)
	log.Println("copied", *repo.Name)
}

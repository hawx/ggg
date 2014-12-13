package main

import (
	"github.com/hawx/ggg/git"
	"github.com/hawx/ggg/repos"

	"code.google.com/p/goauth2/oauth"
	"github.com/google/go-github/github"

	"flag"
	"fmt"
	"log"
)

var (
	dbPath = flag.String("db-path", "", "Path to ggg database")
	gitDir = flag.String("git-dir", "", "Path to ggg git directory")
	token  = flag.String("token", "", "A personal access token")
)

const USAGE = `Usage: ggg-import-github [options]

  Imports all repositories on your GitHub account to ggg.

    --db-path     # Path to ggg database
    --git-dir     # Path to ggg git directory
    --token       # A personal access token
`

func main() {
	flag.Parse()
	if *dbPath == "" || *gitDir == "" || *token == "" {
		fmt.Println(USAGE)
		return
	}

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: *token},
	}

	client := github.NewClient(t.Client())

	db := repos.Open(*dbPath, *gitDir)
	defer db.Close()

	opt := &github.RepositoryListOptions{
		Type: "owner",
	}

	for {
		repos, resp, err := client.Repositories.List("", opt)
		if err != nil {
			log.Fatal(err)
		}

		for _, repo := range repos {
			copyRepo(db, repo)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.ListOptions.Page = resp.NextPage
	}
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

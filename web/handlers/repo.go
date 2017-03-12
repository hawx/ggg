package handlers

import (
	"html/template"
	"path"

	"hawx.me/code/ggg/repos"
	"hawx.me/code/ggg/web/views"

	"hawx.me/code/route"

	"net/http"
)

type Protect func(handler, errHandler http.Handler) http.Handler

func Repo(db repos.Db, title, url string, protect Protect) RepoHandler {
	h := repoHandler{db}

	return RepoHandler{
		Html: h.Html(title, url, protect),
		Git:  h.Git(),
	}
}

type RepoHandler struct {
	Html http.Handler
	Git  http.Handler
}

type repoHandler struct {
	db repos.Db
}

func (h repoHandler) Git() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := route.Vars(r)["name"]
		repoName := name[:len(name)-4]
		repo := h.db.Get(repoName)

		if repo.Name == "" || repo.IsPrivate {
			http.NotFound(w, r)
			return
		}

		http.StripPrefix("/"+name+"/",
			http.FileServer(http.Dir(repo.Path))).ServeHTTP(w, r)
	})
}

func (h repoHandler) Html(title, url string, protect Protect) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := route.Vars(r)["name"]
		repo := h.db.Get(name)

		if repo.Name == "" {
			http.NotFound(w, r)
			return
		}

		innerHandler := h.htmlPage(title, repo, url)

		if repo.IsPrivate {
			protect(innerHandler(true), http.NotFoundHandler()).ServeHTTP(w, r)
			return
		}

		protect(innerHandler(true), innerHandler(false)).ServeHTTP(w, r)
	})
}

func (h repoHandler) htmlPage(title string, repo *repos.Repo, url string) func(bool) http.Handler {
	return func(loggedIn bool) http.Handler {
		r := route.New()

		r.HandleFunc("/:name", func(w http.ResponseWriter, r *http.Request) {
			readmeName, readmeContents := repo.Readme()

			w.Header().Add("Content-Type", "text/html")
			views.Repo.Execute(w, views.RepoCtx{
				Title:        title,
				Url:          url,
				LoggedIn:     loggedIn,
				Name:         repo.Name,
				Web:          repo.Web,
				Description:  repo.Description,
				Path:         repo.Path,
				CloneUrl:     repo.CloneUrl(),
				Files:        repo.Files(""),
				IsEmpty:      repo.IsEmpty(),
				IsPrivate:    repo.IsPrivate,
				ParentDir:    "",
				FileName:     readmeName,
				FileContents: readmeContents,
			})
		})

		r.HandleFunc("/:name/tree/*tree", func(w http.ResponseWriter, r *http.Request) {
			tree := route.Vars(r)["tree"]

			parentDir := path.Dir(tree)

			w.Header().Add("Content-Type", "text/html")
			views.Repo.Execute(w, views.RepoCtx{
				Title:        title,
				Url:          url,
				LoggedIn:     loggedIn,
				Name:         repo.Name,
				Web:          repo.Web,
				Description:  repo.Description,
				Path:         repo.Path,
				CloneUrl:     repo.CloneUrl(),
				Files:        repo.Files(tree),
				IsEmpty:      repo.IsEmpty(),
				IsPrivate:    repo.IsPrivate,
				ParentDir:    parentDir,
				FileName:     "",
				FileContents: "",
			})
		})

		r.HandleFunc("/:name/blob/*blob", func(w http.ResponseWriter, r *http.Request) {
			blob := route.Vars(r)["blob"]

			contents, _ := repo.Contents(blob)
			name := path.Base(blob)

			w.Header().Add("Content-Type", "text/html")
			views.Blob.Execute(w, views.BlobCtx{
				Title:        title,
				Url:          url,
				LoggedIn:     loggedIn,
				Name:         repo.Name,
				Web:          repo.Web,
				Description:  repo.Description,
				Path:         repo.Path,
				CloneUrl:     repo.CloneUrl(),
				IsEmpty:      repo.IsEmpty(),
				IsPrivate:    repo.IsPrivate,
				ParentDir:    "",
				FileName:     name,
				FileContents: template.HTML(contents),
			})
		})

		return r
	}
}

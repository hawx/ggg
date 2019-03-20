package handlers

import (
	"html"
	"html/template"
	"net/http"
	"path"
	"strings"

	"hawx.me/code/ggg/repos"
	"hawx.me/code/route"
)

type Protect func(handler, errHandler http.Handler) http.HandlerFunc

func Repo(db repos.Db, title, url string, protect Protect, templates Templates) RepoHandler {
	h := repoHandler{db, templates}

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
	db        repos.Db
	templates Templates
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

			files, ok := repo.Files("")
			if !ok {
				http.NotFound(w, r)
				return
			}

			w.Header().Add("Content-Type", "text/html")
			h.templates.ExecuteTemplate(w, "repo.gotmpl", RepoCtx{
				Title:        title,
				Url:          url,
				LoggedIn:     loggedIn,
				Name:         repo.Name,
				Web:          repo.Web,
				Description:  repo.Description,
				Path:         repo.Path,
				CloneUrl:     repo.CloneUrl(),
				Files:        files,
				IsEmpty:      repo.IsEmpty(),
				IsPrivate:    repo.IsPrivate,
				Dir:          "",
				ParentDir:    "",
				FileName:     readmeName,
				FileContents: readmeContents,
			})
		})

		r.HandleFunc("/:name/tree/*tree", func(w http.ResponseWriter, r *http.Request) {
			tree := route.Vars(r)["tree"]

			parentDir := path.Dir(tree)
			files, ok := repo.Files(tree)
			if !ok {
				http.NotFound(w, r)
				return
			}

			w.Header().Add("Content-Type", "text/html")
			h.templates.ExecuteTemplate(w, "tree.gotmpl", RepoCtx{
				Title:       title,
				Url:         url,
				LoggedIn:    loggedIn,
				Name:        repo.Name,
				Web:         repo.Web,
				Description: repo.Description,
				Path:        repo.Path,
				CloneUrl:    repo.CloneUrl(),
				Files:       files,
				IsEmpty:     repo.IsEmpty(),
				IsPrivate:   repo.IsPrivate,
				Dir:         tree,
				DirParts:    splitDir(tree),
				ParentDir:   parentDir,
			})
		})

		r.HandleFunc("/:name/blob/*blob", func(w http.ResponseWriter, r *http.Request) {
			blob := route.Vars(r)["blob"]

			contents, ok := repo.Contents(blob)
			if !ok {
				http.NotFound(w, r)
				return
			}

			dir, name := path.Split(blob)
			ext := path.Ext(name)
			if len(ext) > 1 {
				ext = "lang-" + ext[1:]
			} else {
				ext = "nohighlight"
			}

			w.Header().Add("Content-Type", "text/html")
			h.templates.ExecuteTemplate(w, "blob.gotmpl", BlobCtx{
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
				Dir:          prettyDir(dir),
				DirParts:     splitDir(dir),
				ParentDir:    dir,
				FileName:     name,
				FileContents: template.HTML(html.EscapeString(contents)),
				FileLang:     ext,
			})
		})

		return r
	}
}

func splitDir(dir string) (res []PathPart) {
	parts := strings.Split(dir, "/")

	last := ""
	for _, part := range parts {
		if part == "" {
			continue
		}

		last = last + "/" + part
		res = append(res, PathPart{Name: part, Path: last})
	}

	return res
}

func prettyDir(dir string) string {
	if dir == "" {
		return "/"
	}

	return dir[:len(dir)-1]
}

type RepoCtx struct {
	Title    string
	Url      string
	LoggedIn bool

	Name         string
	Web          string
	Description  string
	Path         string
	CloneUrl     string
	Files        []repos.File
	IsEmpty      bool
	IsPrivate    bool
	Dir          string
	DirParts     []PathPart
	ParentDir    string
	FileName     string
	FileContents template.HTML
}

type BlobCtx struct {
	Title    string
	Url      string
	LoggedIn bool

	Name         string
	Web          string
	Description  string
	Path         string
	CloneUrl     string
	IsEmpty      bool
	IsPrivate    bool
	Dir          string
	DirParts     []PathPart
	ParentDir    string
	FileName     string
	FileContents template.HTML
	FileLang     string
}

type PathPart struct {
	Name string
	Path string
}

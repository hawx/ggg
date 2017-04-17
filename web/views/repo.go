package views

import (
	"html/template"

	"hawx.me/code/ggg/repos"
)

const repo = `<!DOCTYPE html>
<html>
  <head>
    {{ template "meta" .Name }}
  </head>
  <body>
    {{ template "header" . }}

    <div class="container">
      {{ template "description" . }}
    </div>

    <aside class="container is-wide">
      {{ if not .IsEmpty }}
        {{ template "files" . }}
      {{ end }}
      {{ template "readme" . }}
    </aside>
  </body>
</html>`

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

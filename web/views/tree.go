package views

import "hawx.me/code/ggg/repos"

const tree = `<!DOCTYPE html>
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
      {{ template "files" . }}
    </aside>
  </body>
</html>`

type TreeCtx struct {
	Title    string
	Url      string
	LoggedIn bool

	Name        string
	Web         string
	Description string
	Path        string
	CloneUrl    string
	Files       []repos.File
	IsEmpty     bool
	IsPrivate   bool
	Dir         string
	DirParts    []PathPart
	ParentDir   string
}

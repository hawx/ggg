package views

import "hawx.me/code/ggg/repos"

const tree = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>{{.Name}}</title>
    <link rel="stylesheet" href="/assets/styles.css"></link>
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
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

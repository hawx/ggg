package views

import (
	"html/template"

	"hawx.me/code/ggg/git"
)

const repo = `<!DOCTYPE html>
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
      <div class="repo {{if .IsPrivate}}private{{end}}">
        <h1><a href="/{{.Name}}">{{.Name}}</a></h1>
        {{if .Web}}&rarr; <a href="{{.Web}}">{{.Web}}</a>{{end}}
        <p>{{.Description}}</p>
        {{ if $.LoggedIn }}
        <div class="buttons">
          <a href="/{{.Name}}/edit">edit</a>
          <a href="/{{.Name}}/delete">delete</a>
        </div>
        {{ end }}

        <pre class="clone"><code>git clone {{.Url}}/{{.CloneUrl}}</code></pre>
      </div>
    </div>

    <aside class="container">
      {{ template "files" . }}
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
	Files        []git.File
	IsEmpty      bool
	IsPrivate    bool
	Dir          string
	DirParts     []PathPart
	ParentDir    string
	FileName     string
	FileContents template.HTML
}

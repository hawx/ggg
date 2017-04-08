package views

import "html/template"

const blob = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>{{.Name}}</title>
    <link rel="stylesheet" href="/assets/styles.css"></link>
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
      {{ template "file" . }}
    </aside>
  </body>
</html>`

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
}

type PathPart struct {
	Name string
	Path string
}

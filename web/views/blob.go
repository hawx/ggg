package views

import "html/template"

const blob = `<!DOCTYPE html>
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
      {{ template "filespath" . }}
      {{ template "file" . }}
    </aside>

    {{ template "highlight" }}
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
	FileLang     string
}

type PathPart struct {
	Name string
	Path string
}

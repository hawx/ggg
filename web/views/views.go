package views

import (
	"html/template"
	"io"

	"hawx.me/code/ggg/web/views/parts"
)

var Blob, Create, Edit, List, Repo, Tree interface {
	Execute(io.Writer, interface{}) error
}

func init() {
	tmpl := template.Must(template.New("noop").Parse(""))

	tmpl = template.Must(tmpl.New("description").Parse(parts.Description))
	tmpl = template.Must(tmpl.New("file").Parse(parts.File))
	tmpl = template.Must(tmpl.New("files").Parse(parts.Files))
	tmpl = template.Must(tmpl.New("filespath").Parse(parts.FilesPath))
	tmpl = template.Must(tmpl.New("header").Parse(parts.Header))
	tmpl = template.Must(tmpl.New("meta").Parse(parts.Meta))
	tmpl = template.Must(tmpl.New("readme").Parse(parts.Readme))

	tmpl = template.Must(tmpl.New("blob").Parse(blob))
	tmpl = template.Must(tmpl.New("create").Parse(create))
	tmpl = template.Must(tmpl.New("edit").Parse(edit))
	tmpl = template.Must(tmpl.New("list").Parse(list))
	tmpl = template.Must(tmpl.New("repo").Parse(repo))
	tmpl = template.Must(tmpl.New("tree").Parse(tree))

	Blob = &wrappedTemplate{tmpl, "blob"}
	Create = &wrappedTemplate{tmpl, "create"}
	Edit = &wrappedTemplate{tmpl, "edit"}
	List = &wrappedTemplate{tmpl, "list"}
	Repo = &wrappedTemplate{tmpl, "repo"}
	Tree = &wrappedTemplate{tmpl, "tree"}
}

type wrappedTemplate struct {
	t *template.Template
	n string
}

func (w *wrappedTemplate) Execute(wr io.Writer, data interface{}) error {
	return w.t.ExecuteTemplate(wr, w.n, data)
}

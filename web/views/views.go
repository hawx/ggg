package views

import (
	"html/template"
	"io"
)

var List, Create, Edit, Repo, Blob interface {
	Execute(io.Writer, interface{}) error
}

func init() {
	tmpl := template.Must(template.New("list").Parse(list))
	tmpl = template.Must(tmpl.New("create").Parse(create))
	tmpl = template.Must(tmpl.New("edit").Parse(edit))
	tmpl = template.Must(tmpl.New("repo").Parse(repo))
	tmpl = template.Must(tmpl.New("header").Parse(header))
	tmpl = template.Must(tmpl.New("files").Parse(files))
	tmpl = template.Must(tmpl.New("readme").Parse(readme))
	tmpl = template.Must(tmpl.New("blob").Parse(blob))

	List = &wrappedTemplate{tmpl, "list"}
	Create = &wrappedTemplate{tmpl, "create"}
	Edit = &wrappedTemplate{tmpl, "edit"}
	Repo = &wrappedTemplate{tmpl, "repo"}
	Blob = &wrappedTemplate{tmpl, "blob"}
}

type wrappedTemplate struct {
	t *template.Template
	n string
}

func (w *wrappedTemplate) Execute(wr io.Writer, data interface{}) error {
	return w.t.ExecuteTemplate(wr, w.n, data)
}

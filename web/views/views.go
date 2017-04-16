package views

import (
	"html/template"
	"io"
)

var List, Create, Edit, Repo, Blob, Tree interface {
	Execute(io.Writer, interface{}) error
}

func init() {
	tmpl := template.Must(template.New("list").Parse(list))
	tmpl = template.Must(tmpl.New("create").Parse(create))
	tmpl = template.Must(tmpl.New("edit").Parse(edit))
	tmpl = template.Must(tmpl.New("repo").Parse(repo))
	tmpl = template.Must(tmpl.New("header").Parse(header))
	tmpl = template.Must(tmpl.New("files").Parse(files))
	tmpl = template.Must(tmpl.New("file").Parse(file))
	tmpl = template.Must(tmpl.New("readme").Parse(readme))
	tmpl = template.Must(tmpl.New("blob").Parse(blob))
	tmpl = template.Must(tmpl.New("tree").Parse(tree))

	List = &wrappedTemplate{tmpl, "list"}
	Create = &wrappedTemplate{tmpl, "create"}
	Edit = &wrappedTemplate{tmpl, "edit"}
	Repo = &wrappedTemplate{tmpl, "repo"}
	Blob = &wrappedTemplate{tmpl, "blob"}
	Tree = &wrappedTemplate{tmpl, "tree"}
}

type wrappedTemplate struct {
	t *template.Template
	n string
}

func (w *wrappedTemplate) Execute(wr io.Writer, data interface{}) error {
	return w.t.ExecuteTemplate(wr, w.n, data)
}

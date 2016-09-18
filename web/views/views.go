package views

import "html/template"

var (
	List   = template.Must(template.New("list").Parse(list))
	Create = template.Must(template.New("create").Parse(create))
	Edit   = template.Must(template.New("edit").Parse(edit))
	Repo   = template.Must(template.New("repo").Parse(repo))
)

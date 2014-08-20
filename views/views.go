package views

import "github.com/hoisie/mustache"

var (
	List,   _ = mustache.ParseString(list)
	Admin,  _ = mustache.ParseString(admin)
	Create, _ = mustache.ParseString(create)
	Edit,   _ = mustache.ParseString(edit)
)

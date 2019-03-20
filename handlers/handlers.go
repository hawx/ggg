package handlers

import "io"

type Templates interface {
	ExecuteTemplate(w io.Writer, name string, data interface{}) error
}

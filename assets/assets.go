package assets

import (
	"net/http"
	"strings"
	"time"
)

func handle(body string) http.Handler {
	return handler{body, time.Now()}
}

type handler struct {
	body string
	at   time.Time
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeContent(w, r, r.URL.Path, h.at, strings.NewReader(h.body))
}

var (
	Styles = handle(styles)
	Core   = handle(core)
)

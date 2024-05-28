package views

import (
	"net/http"
	"strings"
)

// PageHandler !
func PageHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called "/page/"
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/page")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	if strings.HasPrefix(r.URL.Path, "/base") {
		TodoHandler(w, r)
		return
	}
}

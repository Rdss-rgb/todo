package api

import (
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	// ------------------ ADD THIS CODE ------------------
	if strings.HasPrefix(r.URL.Path, "/todo_list") {
		TodoListAPIHandler(w, r)
		return
	}
	// ---------------------------------------------------
	if strings.HasPrefix(r.URL.Path, "/custom_list") {
		CustomListAPIHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/add_friend") {
		AddFriendAPIHandler(w, r)
		return
	}
}

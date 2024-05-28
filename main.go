package main

import (
	// Specify the username that you used inside github.com folder

	"net/http"

	"github.com/Rdss-rgb/todo/api"
	"github.com/Rdss-rgb/todo/models"
	"github.com/Rdss-rgb/todo/views"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.Todo{}, // <-- place it here
		models.Category{},
		models.Friend{},
		models.Items{},
	)
	uadmin.RegisterInlines(models.Category{}, map[string]string{
		"Todo": "CategoryID",
	})
	uadmin.RegisterInlines(models.Friend{}, map[string]string{
		"Todo": "FriendID",
	})
	uadmin.RegisterInlines(models.Items{}, map[string]string{
		"Todo": "ItemsID",
	})
	// API Handler
	http.HandleFunc("/api/", uadmin.Handler(api.Handler))
	http.HandleFunc("/page/", uadmin.Handler(views.PageHandler))

	uadmin.StartServer()
}

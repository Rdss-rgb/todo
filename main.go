package main

import (
	// Specify the username that you used inside github.com folder

	"github.com/Rdss-rgb/todo/models"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.Todo{}, // <-- place it here
	)

	uadmin.StartServer()
}

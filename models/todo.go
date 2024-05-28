package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Todo struct {
	uadmin.Model
	Name        string
	Description string `uadmin:"html"`
	// Category Model
	Category Category
	// Category ID
	CategoryID uint

	// Friend Model
	Friend Friend `uadmin:"help:Who will be a part of your activity?"`
	// Friend ID
	FriendID uint

	// Item Model
	Items Items `uadmin:"help:What are the requirements needed in order to accomplish your activity?"`
	// Item ID
	ItemsID uint

	TargetDate time.Time
	Progress   int `uadmin:"progress_bar"`
}

// Save function !
func (t *Todo) Save() {
	// Save the model to DB
	uadmin.Save(t)
	// Some other business Logic ...
}

// Validate function !
func (t Todo) Validate() (errMsg map[string]string) {
	// Initialize the error messages
	errMsg = map[string]string{}
	// Get any records from the database that maches the name of
	// this record and make sure the record is not the record we are
	// editing right now
	todo := Todo{}
	if uadmin.Count(&todo, "name = ? AND id <> ?", t.Name, t.ID) != 0 {
		errMsg["Name"] = "This todo name is already in the system"
	}
	return
}

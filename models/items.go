package models

import (
	"strings"

	"github.com/uadmin/uadmin"
)

// Item Model !
type Items struct {
	uadmin.Model
	Name        string `uadmin:"required;search;categorical_filter;filter;display_name:Product Name;default_value:Drinks"` // <-- place it here
	Description string `uadmin:"multilingual"`
	// FIELDS ADDED
	Category     []Category `uadmin:"list_exclude" gorm:"many2many:-"`
	CategoryList string     `uadmin:"read_only"`
	Cost         int        `uadmin:"money;pattern:^[0-9]*$;pattern_msg:Your input must be a number.;help:Input numeric characters only in this field."`
	Rating       int        `uadmin:"min:0;max:10"`
}

// Save !
func (i *Items) Save() {
	// Add a new string array type variable called categoryList
	categoryList := []string{}

	// Append every element to the categoryList array
	for c := range i.Category {
		categoryList = append(categoryList, i.Category[c].Name)
	}

	// Concatenate the categoryList to a single string separated by comma
	joinList := strings.Join(categoryList, ", ")

	// Store the joined string to the CategoryList field
	i.CategoryList = joinList

	// Save it to the database
	uadmin.Save(i)
}

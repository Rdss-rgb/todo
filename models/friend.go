package models

import "github.com/uadmin/uadmin"

type Friend struct {
	uadmin.Model
	Name        string `uadmin:"required"`
	Email       string `uadmin:"email"`
	Password    string `uadmin:"password;list_exclude"`
	Nationality Nationality
	Invite      string `uadmin:"link"`
}

type Nationality int

// Chinese !
func (Nationality) Chinese() Nationality {
	return 1
}

// Filipino !
func (Nationality) Filipino() Nationality {
	return 2
}

// Others !
func (Nationality) Others() Nationality {
	return 3
}

func (f *Friend) Save() {
	f.Invite = "https://barekan.netlify.app/"
	uadmin.Save(f)
}

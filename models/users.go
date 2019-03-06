package models

import "github.com/giovannirossini/curso/lib"

// Users struct that maps database to a model
type Users struct {
	ID    int    `db:"id" json:"id"`
	Nome  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

// UsersModel is the table that we'll use in thes CRUD
var UsersModel = lib.Sess.Collection("users")

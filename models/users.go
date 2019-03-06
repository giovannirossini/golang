package models

import "github.com/giovannirossini/curso/lib"

type Users struct {
	ID    int    `db:"id" json:"id"`
	Nome  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
}

var UsersModel = lib.Sess.Collection("users")

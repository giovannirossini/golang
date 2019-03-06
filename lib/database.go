package lib

import (
	"log"

	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var config = mysql.ConnectionURL{
	Host:     "localhost",
	User:     "golang",
	Password: "password",
	Database: "members",
}

// Sess connection var database
var Sess db.Database

func init() {
	var err error

	Sess, err = mysql.Open(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	Sess.SetLogging(true)
}

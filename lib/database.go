package lib

import (
	"log"
    "os"
	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var config = mysql.ConnectionURL{
	Host:     os.Getenv("CLEARDB_HOST"),
	User:     os.Getenv("CLEARDB_USER"),
	Password: os.Getenv("CLEARDB_PASSWD"),
	Database: os.Getenv("CLEARDB_DB"),
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

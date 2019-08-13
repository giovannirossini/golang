package lib

import (
	"log"
    "os"
	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)


var config = mysql.ConnectionURL{
	Host:     "us-cdbr-iron-east-03.cleardb.net",
	User:     "bd51013dec5525",
	Password: os.Getenv("CLEARDB_PASSWD"),
	Database: "heroku_9f89a4257664aca",
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

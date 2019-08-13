package lib

import (
	"log"
    "os"
	"upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var config = os.Getenv("DATABASE_URL")
// var config = mysql.ConnectionURL{
// 	Host:     "us-cdbr-iron-east-03.cleardb.net",
// 	User:     "bd51013dec5525@us-cdbr-iron-east-03.cleardb.net",
// 	Password: "d1be2d4",
// 	Database: "heroku_9f89a4257664aca",
// }

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

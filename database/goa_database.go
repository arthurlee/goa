package database

import (
	"database/sql"
	"github.com/arthurlee/goa/context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var Db *sql.DB = nil

func Init() {
	db, err := sql.Open(context.Instance.Config.Database.Type, context.Instance.Config.Database.Url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Println("ping database ...")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Println("database ok")

	Db = db
}

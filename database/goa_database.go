package database

import (
	"database/sql"
	"github.com/arthurlee/goa/context"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var Db *sql.DB = nil

func init() {
	log.Println("database init")

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

func GetList(dbList DbList) error {
	rows, err := Db.Query(dbList.Sql())
	if err == nil {
		for rows.Next() {
			err := dbList.SetItem(rows)
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		return nil
	}
}

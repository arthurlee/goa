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

func GetList(dbSelect DbSelect) error {
	rows, err := Db.Query(dbSelect.GetSql(), dbSelect.GetArgs()...)
	if err == nil {
		for rows.Next() {
			err := dbSelect.SetItem(rows)
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		return nil
	}
}

func Create(dbOperate DbOperate) (int64, int64, error) {
	result, err := Db.Exec(dbOperate.GetSql(), dbOperate.GetArgs()...)
	if err != nil {
		return 0, 0, err
	}

	lastInsertId, err := result.LastInsertId()
	rowsInserted, err := result.RowsAffected()

	return rowsInserted, lastInsertId, err
}

func Update(dbOperate DbOperate) (int64, error) {
	result, err := Db.Exec(dbOperate.GetSql(), dbOperate.GetArgs()...)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()
	return rowsUpdated, err
}

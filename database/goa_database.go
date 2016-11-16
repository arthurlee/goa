package database

import (
	"database/sql"
	"github.com/arthurlee/goa/instance"
	"github.com/donnie4w/go-logger/logger"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var GoaDatabaseVersion = "0.1.0"

func init() {
	logger.Info("database init")
	_, err := sql.Open(instance.Instance.Config.Database.Type, instance.Instance.Config.Database.Url)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
	logger.Info("database ok")
}

func getDb() (*sql.DB, error) {
	db, err := sql.Open(instance.Instance.Config.Database.Type, instance.Instance.Config.Database.Url)
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
		return nil, err
	}

	return db, nil
}

func GetList(dbSelect DbSelect) error {
	db, err := getDb()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query(dbSelect.GetSql(), dbSelect.GetArgs()...)
	if err == nil {
		for rows.Next() {
			err := dbSelect.SetItem(rows)
			if err != nil {
				return err
			}
		}
		return nil
	} else {
		return err
	}
}

func Create(dbOperate DbOperate) (int64, int64, error) {
	db, err := getDb()
	if err != nil {
		return 0, 0, err
	}
	defer db.Close()

	result, err := db.Exec(dbOperate.GetSql(), dbOperate.GetArgs()...)
	if err != nil {
		return 0, 0, err
	}

	lastInsertId, err := result.LastInsertId()
	rowsInserted, err := result.RowsAffected()

	return rowsInserted, lastInsertId, err
}

func Update(dbOperate DbOperate) (int64, error) {
	db, err := getDb()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	result, err := db.Exec(dbOperate.GetSql(), dbOperate.GetArgs()...)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()
	return rowsUpdated, err
}

func Delete(dbOperate DbOperate) (int64, error) {
	db, err := getDb()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	result, err := db.Exec(dbOperate.GetSql(), dbOperate.GetArgs()...)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := result.RowsAffected()
	return rowsDeleted, err
}

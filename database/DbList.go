package database

import (
	"database/sql"
)

type DbList interface {
	Sql() string
	SetItem(rows *sql.Rows) error
}

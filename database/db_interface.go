package database

import (
	"database/sql"
)

type DbOperate interface {
	Sql() string
}

type DbSelect interface {
	DbOperate
	SetItem(rows *sql.Rows) error
}

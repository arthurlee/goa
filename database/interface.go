package database

import (
	"database/sql"
)

type DbOperate interface {
	GetSql() string
	GetArgs() []interface{}
}

type DbSelect interface {
	DbOperate
	SetItem(rows *sql.Rows) error
}

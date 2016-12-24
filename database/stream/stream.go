package stream

import (
	"bytes"
	"fmt"
	"github.com/arthurlee/goa/database"
	"strings"
)

type dbAction int

const (
	actionSelect = iota
	actionSelectOne
	actionUpdate
	actionDelete
)

const dftMaxRecordCount = 50

// TODO: maybe need a logger instance

type DBStream struct {
	db             *database.Db
	creator        Creator
	data           interface{} // the instance which the creator returns
	action         dbAction
	tableName      string
	fields         []string
	conditions     []string
	maxRecordCount int
}

func Instance(db *database.Db, creator Creator) *DBStream {
	s := DBStream{}

	s.db = db
	s.creator = creator
	s.maxRecordCount = dftMaxRecordCount

	return &s
}

func (me *DBStream) Table(tableName string) *DBStream {
	me.tableName = tableName
	return me
}

func (me *DBStream) Select(fields []string) *DBStream {
	me.action = actionSelect
	me.fields = fields
	return me
}

func (me *DBStream) SelectOne(fields []string) *DBStream {
	me.action = actionSelectOne
	me.fields = fields
	return me
}

func (me *DBStream) Where(conditions []string) *DBStream {
	me.conditions = conditions
	return me
}

func (me *DBStream) genSelect(buffer *bytes.Buffer) {
	buffer.WriteString("SELECT ")
	if me.fields != nil && len(me.fields) > 0 {
		buffer.WriteString(strings.Join(me.fields, ", "))
	} else {
		buffer.WriteString("*")
	}
	buffer.WriteString("\nFROM " + me.tableName)
}

func (me *DBStream) genWhere(buffer *bytes.Buffer) {
	if me.conditions != nil && len(me.conditions) > 0 {
		buffer.WriteString("\nWHERE (" + strings.Join(me.conditions, ") AND (") + ")")
	}
}

func (me *DBStream) Done() (interface{}, error) {
	log := me.db.Log
	log.Debug("Table Name = %s", me.tableName)

	var buffer bytes.Buffer

	if me.action == actionSelectOne || me.action == actionSelect {
		me.genSelect(&buffer)
	}
	me.genWhere(&buffer)
	if me.action == actionSelectOne {
		buffer.WriteString("\nLIMIT 1") // mysql
	}

	sql := buffer.String()
	log.Debug("sql = \n%s", sql)

	record := me.creator()
	err := me.db.Get(record, sql)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(record)
	return record, nil
}

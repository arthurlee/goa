package stream

import (
	"github.com/arthurlee/goa/logger"
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
	creator        Creator
	data           interface{} // the instance which the creator returns
	action         dbAction
	tableName      string
	fields         []string
	conditions     []string
	maxRecordCount int
}

func Instance(creator Creator) *DBStream {
	s := DBStream{}

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

func (me *DBStream) Done() (interface{}, error) {
	logger.Debug("Table Name = %s", me.tableName)
	return nil, nil
}

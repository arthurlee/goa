package database

import (
	"errors"
	"fmt"
	"log"
)

// Database Field type
const (
	DFT_String = iota
	DFT_Number
	DFT_Date     // yyyy-mm-dd
	DFT_Time     // hh:mm:ss
	DFT_DateTime // yyyy-mm-dd hh:mm:ss
)

type DbField struct {
	Name  string
	Type  int
	Value string
}

func (me *DbField) GetSqlValue() string {
	value := ""

	switch me.Type {
	case DFT_String, DFT_Date, DFT_Time, DFT_DateTime:
		value = fmt.Sprintf("'%s'", me.Value)
	default:
		value = me.Value
	}

	return value
}

// ------------------- Operation ---------------------

type DbOperation struct {
	Sql  string
	Args []interface{}
}

func (me DbOperation) GetSql() string {
	return me.Sql
}

func (me DbOperation) GetArgs() []interface{} {
	return me.Args
}

func (me *DbOperation) SetArgs(args ...interface{}) {
	me.Args = args
}

// ------------------- TableName ---------------------

type DbTableName struct {
	TableName string
}

func (me *DbTableName) SetTableName(tableName string) *DbTableName {
	me.TableName = tableName
	return me
}

// ------------------- FieldList ---------------------

type DbFieldList struct {
	Fields []DbField
}

func (me *DbFieldList) AddField(name string, fieldType int, value string) *DbFieldList {
	if me.Fields == nil {
		me.Fields = make([]DbField, 0, 10)
	}

	field := DbField{name, fieldType, value}
	me.Fields = append(me.Fields, field)

	return me
}

func (me *DbFieldList) AddStringField(name string, value string) *DbFieldList {
	return me.AddField(name, DFT_String, value)
}

func (me *DbFieldList) GetFieldNames() string {
	var names = ""
	const SEP = ","
	count := len(me.Fields)
	for i := 0; i < count; i++ {
		names += me.Fields[i].Name
		if i < count-1 {
			names += SEP
		}
	}

	return names
}

func (me *DbFieldList) GetFieldValues() string {
	var values = ""
	const SEP = ","
	count := len(me.Fields)
	for i := 0; i < count; i++ {
		values += me.Fields[i].GetSqlValue()
		if i < count-1 {
			values += SEP
		}
	}

	return values
}

// ------------------- ConditionList ---------------------

type DbConditionList struct {
	Conditions []string
}

func (me *DbConditionList) Where(condition string) *DbConditionList {
	if me.Conditions == nil {
		me.Conditions = make([]string, 0, 10)
	}
	me.Conditions = append(me.Conditions, condition)
	return me
}

// ------------------- Inserter ---------------------

type DbInserter struct {
	DbOperation
	DbTableName
	DbFieldList
	DbConditionList
}

func (me *DbInserter) Done() error {
	// Sql can be provided manually, it is ok
	if len(me.Sql) > 0 {
		return nil
	}

	if len(me.TableName) == 0 {
		return errors.New("please provide the table name")
	}

	if me.Fields == nil {
		return errors.New("please provide the fields to insert")
	}

	me.Sql = fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s)", me.TableName, me.GetFieldNames(), me.GetFieldValues())

	log.Printf("sql: %q", me.Sql)

	return nil
}

// ------------------- Updater ---------------------

type DbUpdate interface {
}

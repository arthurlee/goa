package checkparam

import (
	"errors"
	"fmt"
	"github.com/arthurlee/goa/server"
)

//
// check if the parameter is exist, is specific length
//

type IRange interface {
	GetMin() int64
	GetMax() int64
}

type IntegerRangeCheckItem struct {
	RegexCheckItem
	min int64
	max int64
}

func (me *IntegerRangeCheckItem) GetMin() int64 {
	return me.min
}

func (me *IntegerRangeCheckItem) GetMax() int64 {
	return me.max
}

func HandlerIntegerRange(item CheckBase, ctx *server.HttpContext) (interface{}, error) {
	val, err := HandlerInteger(item, ctx)
	if err != nil {
		return val, err
	}

	iVal := val.(int64)
	intRange := item.(IRange)
	if iVal < intRange.GetMin() || iVal > intRange.GetMax() {
		return nil, errors.New(fmt.Sprintf("parameter %s is not between %d and %d !", item.GetName(), intRange.GetMin(), intRange.GetMax()))
	}

	return val, err
}

func IntegerRange(name string, min int64, max int64, errorCode string) CheckBase {
	return &IntegerRangeCheckItem{RegexCheckItem{CheckItem{name, errorCode, HandlerIntegerRange}, "^\\d+$"}, min, max}
}

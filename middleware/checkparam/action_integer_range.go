package checkparam

import (
	"errors"
	"fmt"
	"github.com/arthurlee/goa/server"
)

//
// check if the parameter is integer and in the range of min and max
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
	intRange := item.(IRange)
	val, err := HandlerInteger(item, ctx)
	if err != nil {
		return val, errors.New(fmt.Sprintf("parameter %s is not between %d and %d !", item.GetName(), intRange.GetMin(), intRange.GetMax()))
	}

	_, ok := val.(string)
	if !item.IsRequired() && ok {
		return val, nil
	}

	iVal := val.(int64)
	if iVal < intRange.GetMin() || iVal > intRange.GetMax() {
		return nil, errors.New(fmt.Sprintf("parameter %s is not between %d and %d !", item.GetName(), intRange.GetMin(), intRange.GetMax()))
	}

	return val, err
}

func IntegerRange(name string, min int64, max int64, errorCode string, required bool) CheckBase {
	return &IntegerRangeCheckItem{RegexCheckItem{CheckItem{name, errorCode, required, HandlerIntegerRange}, "^\\d+$"}, min, max}
}

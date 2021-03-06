package checkparam

import (
	"github.com/arthurlee/goa/server"
	"strconv"
)

//
// check if the parameter is integer
//

type IntegerCheckItem struct {
	RegexCheckItem
}

func HandlerInteger(item CheckBase, ctx *server.HttpContext) (interface{}, error) {
	val, err := HandlerRegex(item, ctx)
	if err != nil {
		return val, err
	}

	iVal, err := strconv.ParseInt(val.(string), 10, 64)
	return iVal, err
}

func Integer(name string, errorCode string, required bool) CheckBase {
	return &RegexCheckItem{CheckItem{name, errorCode, required, HandlerInteger}, "^\\d+$"}
}

package checkparam

import (
	"github.com/arthurlee/goa/server"
	"strconv"
)

//
// check if the parameter is exist, is specific length
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

func Integer(name string, errorCode string) CheckBase {
	return &RegexCheckItem{CheckItem{name, errorCode, HandlerInteger}, "^\\d+$"}
}

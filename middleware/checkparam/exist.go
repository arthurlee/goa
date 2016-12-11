package checkparam

import (
	"errors"
	"fmt"
	"github.com/arthurlee/goa/server"
	"strings"
)

//
// check if the parameter is exist, is present and not empty
//

func HandlerExist(item *CheckItem, ctx *server.HttpContext) error {
	val := strings.Join(ctx.Form[item.Name], "")
	ctx.Log.Debug("HandlerExist: %s = %s", item.Name, val)
	if len(val) == 0 {
		return errors.New(fmt.Sprintf("parameter %s is empty", item.Name))
	}

	return nil
}

func Exist(name string, errorCode string) *CheckItem {
	return &CheckItem{name, errorCode, HandlerExist}
}

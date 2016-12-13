package checkparam

import (
	"errors"
	"fmt"
	"github.com/arthurlee/goa/server"
	//"strings"
)

//
// check if the parameter is exist, is present and not empty
//

func HandlerExist(item CheckBase, ctx *server.HttpContext) error {
	name := item.GetName()

	val := ctx.R.Form.Get(name)
	ctx.Log.Debug("HandlerExist: %s = '%s'", name, val)
	if len(val) == 0 {
		return errors.New(fmt.Sprintf("parameter %s is empty", name))
	}

	return nil
}

func Exist(name string, errorCode string) CheckBase {
	return &CheckItem{name, errorCode, HandlerExist}
}
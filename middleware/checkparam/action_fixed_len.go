package checkparam

import (
	"errors"
	"fmt"
	"github.com/arthurlee/goa/server"
	//"strings"
)

//
// check if the parameter is exist, is specific length
//

type IFixedLen interface {
	GetFixedLen() int
}

type FixedLenCheckItem struct {
	CheckItem
	fixedLen int
}

func (me *FixedLenCheckItem) GetFixedLen() int {
	return me.fixedLen
}

func HandlerFixedLen(item CheckBase, ctx *server.HttpContext) (interface{}, error) {
	name := item.GetName()

	//val := strings.Join(ctx.Form[name], "")
	val := ctx.R.Form.Get(name)
	ctx.Log.Debug("HandlerFixedLen: %s = %s", name, val)

	fixedLen := item.(IFixedLen)
	if len(val) != fixedLen.GetFixedLen() {
		return nil, errors.New(fmt.Sprintf("parameter %s's length is not %d", name, fixedLen.GetFixedLen()))
	}

	return val, nil
}

func FixedLen(name string, length int, errorCode string) CheckBase {
	return &FixedLenCheckItem{CheckItem{name, errorCode, HandlerFixedLen}, length}
}

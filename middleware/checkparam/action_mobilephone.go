package checkparam

import (
	"errors"
	"fmt"
	"github.com/arthurlee/goa/server"
)

//
// check if the parameter is mobilephone
//

type MobilephoneCheckItem struct {
	RegexCheckItem
}

func HandlerMobilephone(item CheckBase, ctx *server.HttpContext) (interface{}, error) {
	val, err := HandlerRegex(item, ctx)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Parameter %s is not a mobilephone number!", item.GetName()))
	}

	return val, nil
}

func Mobilephone(name string, errorCode string) CheckBase {
	return &RegexCheckItem{CheckItem{name, errorCode, HandlerMobilephone}, "^\\d{11}$"}
}

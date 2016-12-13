package checkparam

import (
	"errors"
	"fmt"
	"github.com/arthurlee/goa/server"
	//"strings"
	"regexp"
)

//
// check if the parameter is exist, is specific length
//

type IRegex interface {
	IsMatch(value string) bool
	GetPattern() string
}

type RegexCheckItem struct {
	CheckItem
	pattern string
}

func (me *RegexCheckItem) IsMatch(value string) bool {
	m, _ := regexp.MatchString(me.pattern, value)
	return m
}

func (me *RegexCheckItem) GetPattern() string {
	return me.pattern
}

func HandlerRegex(item CheckBase, ctx *server.HttpContext) error {
	name := item.GetName()

	val := ctx.R.Form.Get(name)
	ctx.Log.Debug("HandlerRegex: %s = %s", name, val)

	regex := item.(IRegex)
	if !regex.IsMatch(val) {
		return errors.New(fmt.Sprintf("parameter %s does match the pattern '%s' !", name, regex.GetPattern()))
	}

	return nil
}

func Regex(name string, pattern string, errorCode string) CheckBase {
	return &RegexCheckItem{CheckItem{name, errorCode, HandlerRegex}, pattern}
}

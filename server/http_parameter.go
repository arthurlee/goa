package server

import (
//"strconv"
)

type HttpParam map[string]interface{}

func HttpParamCreate() HttpParam {
	m := make(map[string]interface{})
	return m
}

func (me *HttpParam) Set(key string, value interface{}) {
	(*me)[key] = value
}

func (me *HttpParam) Get(key string) (interface{}, bool) {
	v, ok := (*me)[key]
	return v, ok
}

func (me *HttpParam) GetStr(key string) (string, bool) {
	v, ok := me.Get(key)
	if !ok {
		return "", ok
	}

	val, ok := v.(string)
	return val, ok
}

func (me *HttpParam) GetInt(key string) (int, bool) {
	v, ok := me.Get(key)
	if !ok {
		return 0, ok
	}

	val, ok := v.(int)
	return val, ok
}

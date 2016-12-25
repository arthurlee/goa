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

// get

func (me *HttpParam) Get(key string) (interface{}, bool) {
	v, ok := (*me)[key]
	return v, ok
}

func (me *HttpParam) GetWithDft(key string, dft interface{}) (interface{}, bool) {
	v, ok := (*me)[key]
	if !ok {
		return dft, true
	}

	return v, true
}

// string

func (me *HttpParam) GetStr(key string) (string, bool) {
	v, ok := me.Get(key)
	if !ok {
		return "", ok
	}

	val, ok := v.(string)
	return val, ok
}

func (me *HttpParam) GetStrWithDft(key string, dft string) (string, bool) {
	v, ok := me.Get(key)
	if !ok {
		return dft, true
	}

	val, ok := v.(string)
	return val, ok
}

// int

func (me *HttpParam) GetInt(key string) (int, bool) {
	v, ok := me.Get(key)
	if !ok {
		return 0, ok
	}

	val, ok := v.(int)
	return val, ok
}

func (me *HttpParam) GetIntWithDft(key string, dft int) (int, bool) {
	v, ok := me.Get(key)
	if !ok {
		return dft, true
	}

	val, ok := v.(int)
	return val, ok
}

// int64

func (me *HttpParam) GetInt64(key string) (int64, bool) {
	v, ok := me.Get(key)
	if !ok {
		return 0, ok
	}

	val, ok := v.(int64)
	return val, ok
}

func (me *HttpParam) GetInt64WithDft(key string, dft int64) (int64, bool) {
	v, ok := me.Get(key)
	if !ok {
		return dft, true
	}

	val, ok := v.(int64)
	return val, ok
}

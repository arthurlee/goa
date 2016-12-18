package server

import (
	"errors"
	"testing"
)

// go test

func TestCreate(t *testing.T) {
	m := HttpParamCreate()
	if m == nil {
		t.Error(errors.New("create error"))
	}
}

func TestGetSet(t *testing.T) {
	m := HttpParamCreate()

	m.Set("name", "hello")
	val, ok := m.Get("name")
	if !ok {
		t.Error(errors.New("cannot get the value of the key"))
		return
	}

	strVal, ok := val.(string)
	if !ok {
		t.Error(errors.New("value is not string"))
		return
	}

	if strVal != "hello" {
		t.Error(errors.New("value is not correct"))
		return
	}
}

func TestGetStr(t *testing.T) {
	m := HttpParamCreate()

	m.Set("name", "hello")
	val, ok := m.GetStr("name")
	if !ok {
		t.Error(errors.New("cannot get the value of the key"))
		return
	}

	if val != "hello" {
		t.Error(errors.New("value is not correct"))
		return
	}
}

func TestGetInt(t *testing.T) {
	m := HttpParamCreate()

	m.Set("name", 123)
	val, ok := m.GetInt("name")
	if !ok {
		t.Error(errors.New("cannot get the value of the key"))
		return
	}

	if val != 123 {
		t.Error(errors.New("value is not correct"))
		return
	}
}

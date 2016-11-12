package server

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type GoaResponse struct {
	W *http.ResponseWriter
	R *http.Request

	Form url.Values
}

func (me *GoaResponse) SendJson(data interface{}) {
	err := json.NewEncoder(*me.W).Encode(data)
	if err != nil {
		log.Println(err)
	}
}

type GoaBaseRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (me *GoaResponse) SendError(code string, message string) {
	res := GoaBaseRes{code, message}
	me.SendJson(res)
}

func (me *GoaResponse) SendOK() {
	res := GoaBaseRes{"0", "ok"}
	me.SendJson(res)
}

// ------------------------------------
// convenient functions

func GetSendOK() func(*GoaResponse) {
	return func(res *GoaResponse) {
		res.SendOK()
	}
}

func GetSendError(code string, message string) func(*GoaResponse) {
	return func(res *GoaResponse) {
		res.SendError(code, message)
	}
}

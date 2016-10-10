package route

import (
	"encoding/json"
	"net/http"
)

type GoaHandler func(*GoaResponse)

func Get(path string, handler GoaHandler) {
	getHandlerMap[path] = handler
}

func Post(path string, handler GoaHandler) {
	postHandlerMap[path] = handler
}

type GoaResponse struct {
	w *http.ResponseWriter
	r *http.Request
}

func (goaRes *GoaResponse) SendJson(data interface{}) {
	json.NewEncoder(*goaRes.w).Encode(data)
}

type GoaBaseRes struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (goaRes *GoaResponse) SendError(code string, message string) {
	res := GoaBaseRes{code, message}
	goaRes.SendJson(res)
}

func (goaRes *GoaResponse) SendOK() {
	res := GoaBaseRes{"0", "ok"}
	goaRes.SendJson(res)
}

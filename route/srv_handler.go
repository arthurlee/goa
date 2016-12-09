package route

import (
	// "encoding/json"
	"github.com/arthurlee/goa/middleware"
	"github.com/arthurlee/goa/server"
	"net/http"
)

type tSrvHandler struct{}

var SrvHandler tSrvHandler

func (tSrvHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	context := server.CreateHttpContext(w, r)

	for e := MiddlewareList.Front(); e != nil; e = e.Next() {
		m := e.Value.(*middleware.Entry)
		context.Log.Debug("Do middleware [%s] ...", m.Name)
		m.Handler(context)
	}

	// w.Header().Set("Content-Type", "application/json")
	//
	// handler, err := getHandler(r.Method, r.URL.Path)
	// if err != nil {
	// 	sendError(err, w)
	// 	return
	// }
	//
	// r.ParseForm()
	//
	// context := server.HttpContext{W: w, R: r, Form: r.Form}
	// handler(&context)
}

// func sendError(err error, w http.ResponseWriter) {
// 	errRes := struct {
// 		Code    string `json:"code"`
// 		Message string `json:"message"`
// 	}{"-1", err.Error()}
// 	json.NewEncoder(w).Encode(errRes)
// }

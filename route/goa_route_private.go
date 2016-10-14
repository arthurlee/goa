package route

import (
	"encoding/json"
	"errors"
	"github.com/arthurlee/goa/server"
	"log"
	"net/http"
)

type tHandlerMap map[string]GoaHandler

var getHandlerMap, postHandlerMap tHandlerMap

func init() {
	//log.Println("Goa route init")

	getHandlerMap = tHandlerMap{}
	postHandlerMap = tHandlerMap{}

	//http.HandleFunc("*", goaRouteHandle)
}

type tSrvHandler struct{}

var SrvHandler tSrvHandler

func (a tSrvHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	log.Println(r.Method, r.URL.Path)

	handler, err := getHandler(r.Method, r.URL.Path)
	if err != nil {
		sendError(err, w)
		return
	}

	r.ParseForm()

	goaRes := server.GoaResponse{&w, r, r.Form}
	handler(&goaRes)
}

func sendError(err error, w http.ResponseWriter) {
	errRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{"-1", err.Error()}
	json.NewEncoder(w).Encode(errRes)
}

func getHandler(method string, path string) (GoaHandler, error) {
	var handler GoaHandler = nil
	var err error = nil

	switch method {
	case "GET":
		handler, err = getHandlerFromMap(getHandlerMap, path)
	case "POST":
		handler, err = getHandlerFromMap(postHandlerMap, path)
	default:
		err = errors.New("method does not support")
	}

	return handler, err
}

func getHandlerFromMap(handlerMap tHandlerMap, path string) (GoaHandler, error) {
	handler, ok := handlerMap[path]
	if ok {
		return handler, nil
	} else {
		return nil, errors.New("url does not support")
	}
}

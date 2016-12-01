package route

import (
	"container/list"
	"encoding/json"
	"errors"
	"github.com/arthurlee/goa/logger"
	"github.com/arthurlee/goa/middleware"
	"github.com/arthurlee/goa/server"
	"net/http"
)

type tHandlerMap map[string]server.HttpHandler

var getHandlerMap, postHandlerMap tHandlerMap

func init() {
	getHandlerMap = tHandlerMap{}
	postHandlerMap = tHandlerMap{}

	//http.HandleFunc("*", goaRouteHandle)
}

type tSrvHandler struct{}

var SrvHandler tSrvHandler

func (tSrvHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	logger.Info("%s %s", r.Method, r.URL.Path)

	handler, err := getHandler(r.Method, r.URL.Path)
	if err != nil {
		sendError(err, w)
		return
	}

	r.ParseForm()

	context := server.HttpContext{W: &w, R: r, Form: r.Form}
	handler(&context)
}

func sendError(err error, w http.ResponseWriter) {
	errRes := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{"-1", err.Error()}
	json.NewEncoder(w).Encode(errRes)
}

func getHandler(method string, path string) (server.HttpHandler, error) {
	var handler server.HttpHandler = nil
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

func getHandlerFromMap(handlerMap tHandlerMap, path string) (server.HttpHandler, error) {
	handler, ok := handlerMap[path]
	if ok {
		return handler, nil
	} else {
		return nil, errors.New("url does not support")
	}
}

// middleware support

var middlewareList = list.New()

func routeUse(rm *middleware.Entry, after string) {
	var element *list.Element = nil
	if len(after) > 0 {
		for e := middlewareList.Front(); e != nil; e = e.Next() {
			if e.Value.(*middleware.Entry).Name == rm.Name {
				element = e
				break
			}
		}
	}

	if element != nil {
		middlewareList.InsertAfter(rm, element)
	} else {
		middlewareList.PushBack(rm)
	}
}

func routeDumpRoutes() {
	logger.Info("------------------- route list -------------------")
	for e := middlewareList.Front(); e != nil; e = e.Next() {
		m := e.Value.(*middleware.Entry)
		logger.Info("    %20s    %s", m.Name, m.Version)
	}
	logger.Info("--------------------------------------------------")
}

package route

import (
// "container/list"
// "encoding/json"
// "errors"
// "github.com/arthurlee/goa/logger"
// "github.com/arthurlee/goa/middleware"
// "github.com/arthurlee/goa/server"
// "net/http"
)

// type tHandlerMap map[string]server.HttpHandler
//
// var getHandlerMap, postHandlerMap tHandlerMap
//
// func init() {
// 	getHandlerMap = tHandlerMap{}
// 	postHandlerMap = tHandlerMap{}
//
// 	//http.HandleFunc("*", goaRouteHandle)
// }
//
// func getHandler(method string, path string) (server.HttpHandler, error) {
// 	var handler server.HttpHandler = nil
// 	var err error = nil
//
// 	switch method {
// 	case "GET":
// 		handler, err = getHandlerFromMap(getHandlerMap, path)
// 	case "POST":
// 		handler, err = getHandlerFromMap(postHandlerMap, path)
// 	default:
// 		err = errors.New("method does not support")
// 	}
//
// 	return handler, err
// }
//
// func getHandlerFromMap(handlerMap tHandlerMap, path string) (server.HttpHandler, error) {
// 	handler, ok := handlerMap[path]
// 	if ok {
// 		return handler, nil
// 	} else {
// 		return nil, errors.New("url does not support")
// 	}
// }

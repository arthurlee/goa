package handler

import (
	"errors"
	"github.com/arthurlee/goa/middleware"
	"github.com/arthurlee/goa/server"
)

// TODO: use another map for http methods

type tHandlerMap map[string]server.HttpHandler

var getHandlerMap, postHandlerMap tHandlerMap

func HttpGet(path string, handler server.HttpHandler) {
	getHandlerMap[path] = handler
}

func HttpPost(path string, handler server.HttpHandler) {
	postHandlerMap[path] = handler
}

func init() {
	getHandlerMap = tHandlerMap{}
	postHandlerMap = tHandlerMap{}
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

// middleware entry

var RM_httpHandle = middleware.Entry{"HttpHandle", "0.0.1", httpHandle}

func httpHandle(ctx *server.HttpContext) (server.HResult, error) {

	handler, err := getHandler(ctx.R.Method, ctx.R.URL.Path)
	if err != nil {
		ctx.SendError("-1", err.Error())
		return server.HR_ERROR, err
	}

	return handler(ctx)
}

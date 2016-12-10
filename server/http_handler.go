package server

import (
	"errors"
	"github.com/arthurlee/goa/logger"
)

type HResult int

const (
	HR_OK   HResult = iota
	HR_WARN         // can continue
	HR_ERROR
)

type HttpHandler func(*HttpContext) (HResult, error)

type HttpHandlerItem struct {
	Handler HttpHandler
}

type tHandlerMap map[string]HttpHandlerItem
type tMethodMap map[string]tHandlerMap

var methodMap tMethodMap

func HttpGet(path string, handler HttpHandler) {
	handlerMap := getOrCreateMethodHandlerMap("GET")
	handlerMap[path] = HttpHandlerItem{handler}
}

func HttpPost(path string, handler HttpHandler) {
	handlerMap := getOrCreateMethodHandlerMap("POST")
	handlerMap[path] = HttpHandlerItem{handler}
}

func DumpHttpHandlers() {
	logger.Info("------------------- http handler list -------------------")
	for kMethod, kHandlerMap := range methodMap {
		for kPath, _ := range kHandlerMap {
			logger.Info("    %-10s %s", kMethod, kPath)
		}
	}
	logger.Info("--------------------------------------------------")
}

func init() {
	methodMap = tMethodMap{}
}

func getOrCreateMethodHandlerMap(method string) tHandlerMap {
	handlerMap, ok := methodMap[method]
	if !ok {
		handlerMap = tHandlerMap{}
		methodMap[method] = handlerMap
		logger.Debug("Create http method map for %s", method)
	}
	return handlerMap
}

func GetHandlerItem(method string, path string) (HttpHandlerItem, error) {
	var item HttpHandlerItem = HttpHandlerItem{nil}
	var err error = nil

	handlerMap, ok := methodMap[method]
	if ok {
		item, err = getHandlerFromMap(handlerMap, path)
	} else {
		err = errors.New("method does not support")
	}

	return item, err
}

func getHandlerFromMap(handlerMap tHandlerMap, path string) (HttpHandlerItem, error) {
	item, ok := handlerMap[path]
	if ok {
		return item, nil
	} else {
		return HttpHandlerItem{nil}, errors.New("url does not support")
	}
}

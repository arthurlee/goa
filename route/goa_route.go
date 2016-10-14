package route

import (
	"github.com/arthurlee/goa/server"
)

type GoaHandler func(*server.GoaResponse)

func Get(path string, handler GoaHandler) {
	getHandlerMap[path] = handler
}

func Post(path string, handler GoaHandler) {
	postHandlerMap[path] = handler
}

package route

import (
	"github.com/arthurlee/goa/middleware"
	"github.com/arthurlee/goa/server"
)

// middleware support

func Use(rm *middleware.Entry, after string) {
	routeUse(rm, after)
}

func DumpRoutes() {
	routeDumpRoutes()
}

func Get(path string, handler server.HttpHandler) {
	getHandlerMap[path] = handler
}

func Post(path string, handler server.HttpHandler) {
	postHandlerMap[path] = handler
}

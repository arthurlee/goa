package route

import (
	"github.com/arthurlee/goa/middleware"
	"github.com/arthurlee/goa/server"
)

// convenient purpose

func Use(rm *middleware.Entry) {
	UseAfter(rm, "")
}

func Get(path string, h server.HttpHandler) {
	server.HttpGet(path, h)
}

func Post(path string, h server.HttpHandler) {
	server.HttpPost(path, h)
}

func DumpHttpHandlers() {
	server.DumpHttpHandlers()
}

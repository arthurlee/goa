package route

import (
	"github.com/arthurlee/goa/middleware/handler"
	"github.com/arthurlee/goa/server"
)

func Get(path string, h server.HttpHandler) {
	handler.HttpGet(path, h)
}

func Post(path string, h server.HttpHandler) {
	handler.HttpPost(path, h)
}

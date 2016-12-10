package middleware

import (
	"github.com/arthurlee/goa/server"
)

var RM_DoHttpHandler = Entry{"DoHttpHandler", "0.0.1", doHttpHandler}

func doHttpHandler(ctx *server.HttpContext) (server.HResult, error) {
	val, _ := ctx.Get("handler_item")
	handlerItem := val.(server.HttpHandlerItem)
	return handlerItem.Handler(ctx)
}

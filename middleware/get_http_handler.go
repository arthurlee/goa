package middleware

import (
	"github.com/arthurlee/goa/server"
)

var RM_GetHttpHandler = Entry{"GetHttpHandler", "0.0.1", getHttpHandler}

func getHttpHandler(ctx *server.HttpContext) (server.HResult, error) {

	handlerItem, err := server.GetHandlerItem(ctx.R.Method, ctx.R.URL.Path)
	if err != nil {
		return server.HR_ERROR, err
	}

	ctx.Set("handler_item", handlerItem)
	return server.HR_OK, nil
}

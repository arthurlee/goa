package handler

import (
	"github.com/arthurlee/goa/middleware"
	"github.com/arthurlee/goa/server"
)

var RM_jsonPrepare = middleware.Entry{"JsonPrepare", "0.0.1", jsonPrepare}

func jsonPrepare(ctx *server.HttpContext) (server.HResult, error) {
	ctx.W.Header().Set("Content-Type", "application/json")
	ctx.R.ParseForm()

	return server.HR_OK, nil
}

package middleware

import (
	"github.com/arthurlee/goa/middleware"
	"github.com/arthurlee/goa/server"
	"time"
)

var RM_logStart = middleware.Entry{"LogStart", "0.0.1", logStart}

func logStart(ctx *server.HttpContext) (server.HResult, error) {
	ctx.Set("start_time", time.Now())
	return server.HR_OK, nil
}

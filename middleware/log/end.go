package middleware

import (
	"github.com/arthurlee/goa/middleware"
	"github.com/arthurlee/goa/server"
	"time"
)

var RM_logEnd = middleware.Entry{"LogEnd", "0.0.1", logEnd}

func logEnd(ctx *server.HttpContext) (server.HResult, error) {
	ctx.Set("end_time", time.Now())
	return server.HR_OK, nil
}

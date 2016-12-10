package middleware

import (
	"github.com/arthurlee/goa/server"
	"time"
)

var RM_StartLog = Entry{"StartLog", "0.0.1", startLog}

func startLog(ctx *server.HttpContext) (server.HResult, error) {
	ctx.Set("start_time", time.Now())

	ctx.Log.Info("[request] %s %s", ctx.R.Method, ctx.R.URL.Path)

	return server.HR_OK, nil
}

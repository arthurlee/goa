package middleware

import (
	"github.com/arthurlee/goa/server"
	"time"
)

var RM_EndLog = Entry{"EndLog", "0.0.1", endLog}

func endLog(ctx *server.HttpContext) (server.HResult, error) {
	endTime := time.Now()
	ctx.Set("end_time", endTime)

	val, _ := ctx.Get("start_time")
	startTime := val.(time.Time)
	ms := endTime.Sub(startTime) / time.Millisecond

	ctx.Log.Info("[response] %s %s %dms", ctx.R.Method, ctx.R.URL.Path, ms)

	return server.HR_OK, nil
}

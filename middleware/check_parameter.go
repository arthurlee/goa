package middleware

import (
	"github.com/arthurlee/goa/server"
)

var RM_CheckParameter = Entry{"CheckParameter", "0.0.1", checkParameter}

func checkParameter(ctx *server.HttpContext) (server.HResult, error) {
	return server.HR_OK, nil
}

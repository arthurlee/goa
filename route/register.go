package route

import (
	"github.com/arthurlee/goa/middleware/handler"
	MLog "github.com/arthurlee/goa/middleware/log"
)

// TODO: add custom callback to register

func Register() {
	Use(&MLog.RM_logStart)
	Use(&handler.RM_jsonPrepare)
	Use(&handler.RM_httpHandle)
	Use(&MLog.RM_logEnd)

	DumpRoutes()
}

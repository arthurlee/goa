package route

import (
	MLog "github.com/arthurlee/goa/middleware/log"
)

func Register() {
	Use(&MLog.RM_logStart, "")
	Use(&MLog.RM_logEnd, "")

	DumpRoutes()
}

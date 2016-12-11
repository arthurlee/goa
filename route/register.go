package route

import (
	"github.com/arthurlee/goa/middleware"
	"github.com/arthurlee/goa/middleware/checkparam"
)

// TODO: add custom callback to register
// TODO: let logStart and logEnd always execute

func Register() {
	Use(&middleware.RM_StartLog)
	Use(&middleware.RM_JsonPrepare)
	Use(&middleware.RM_GetHttpHandler)
	Use(&checkparam.RM_CheckParameter)
	Use(&middleware.RM_DoHttpHandler)
	Use(&middleware.RM_EndLog)

	DumpMiddlewares()
}

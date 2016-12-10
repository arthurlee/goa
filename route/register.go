package route

import ()

// TODO: add custom callback to register
// TODO: let logStart and logEnd always execute

func Register() {
	Use(&RM_StartLog)
	Use(&RM_JsonPrepare)
	Use(&RM_GetHttpHandler)
	Use(&RM_CheckParameter)
	Use(&RM_DoHttpHandler)
	Use(&RM_EndLog)

	DumpMiddlewares()
}

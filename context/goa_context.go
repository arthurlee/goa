package context

type GoaContext struct {
	AppRootPath       string
	AppConfigFilePath string
	Config            GoaConfig
}

var Instance GoaContext

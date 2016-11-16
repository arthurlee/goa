package instance

type GoaInstance struct {
	AppRootPath       string
	AppConfigFilePath string
	Config            GoaConfig
}

var Instance GoaInstance

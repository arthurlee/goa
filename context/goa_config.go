package context

type GoaConfigServer struct {
	Address string `yaml:"address"`
}

type GoaConfigDatabase struct {
	Type string `yaml:"type"`
	Url  string `yaml:"url"`
}

type GoaConfig struct {
	Server   GoaConfigServer   `yaml:"server"`
	Database GoaConfigDatabase `yaml:"database"`
}

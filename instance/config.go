package instance

type GoaConfigServer struct {
	Address      string `yaml:"address"`
	HttpsAddress string `yaml:"https_address"`
	CertFile     string `yaml:"cert_file"`
	KeyFile      string `yaml:"key_file"`
}

type GoaConfigDatabase struct {
	Type string `yaml:"type"`
	Url  string `yaml:"url"`
}

type GoaConfigLogger struct {
	Console  bool   `yaml:"console"`
	Dir      string `yaml:"dir"`
	Filename string `yaml:"filename"`
	Level    string `yaml:"level"`
}

type GoaConfig struct {
	Server   GoaConfigServer   `yaml:"server"`
	Database GoaConfigDatabase `yaml:"database"`
	Logger   GoaConfigLogger   `yaml:"logger"`
}

package instance

import (
	"errors"
	"github.com/arthurlee/goa/file"
	"github.com/donnie4w/go-logger/logger"
	"github.com/kardianos/osext"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

func init() {
	folder, err := guessAppRootPath()
	if err != nil {
		logger.Fatal("Goa: ", err)
		os.Exit(1)
	}

	Instance.setAppRoot(folder)

	_, err = Instance.loadConfig()
	if err != nil {
		logger.Fatal("Goa: ", err)
		os.Exit(1)
	}
}

func (me *GoaInstance) setAppRoot(appRoot string) {
	me.AppRootPath = appRoot
	me.AppConfigFilePath = path.Join(appRoot, "conf", "goa.yaml")
}

func (me *GoaInstance) loadConfig() (bool, error) {
	yamlContent, err := ioutil.ReadFile(me.AppConfigFilePath)
	if err != nil {
		return false, err
	}

	err = yaml.Unmarshal(yamlContent, &me.Config)
	if err != nil {
		return false, err
	}

	me.initLogger()

	return true, nil
}

func (me *GoaInstance) initLogger() {
	logConfig := me.Config.Logger
	logger.SetConsole(logConfig.Console)

	level := logger.INFO

	switch logConfig.Level {
	case "all":
		level = logger.ALL
	case "debug":
		level = logger.DEBUG
	case "info":
		level = logger.INFO
	case "warn":
		level = logger.WARN
	case "error":
		level = logger.ERROR
	case "fatal":
		level = logger.FATAL
	case "off":
		level = logger.OFF
	}

	logger.SetLevel(level)

	if logConfig.Dir == "" {
		logConfig.Dir = "log"
	}
	if logConfig.Filename == "" {
		logConfig.Filename = "app.log"
	}

	logger.SetRollingDaily(logConfig.Dir, logConfig.Filename)

	logger.Info("------------------------- start -----------------------")
}

// check if goa.yaml exists
func isGoaApp(folder string) bool {
	cfgFilePath := path.Join(folder, "conf", "goa.yaml")
	return fileutil.IsFile(cfgFilePath)
}

func checkAppRoot(folder string, err error) (string, error) {
	if err == nil {
		if isGoaApp(folder) {
			return folder, nil
		}
		return "", errors.New("Cannot detect the goa app!")
	} else {
		// it rarely happened
		logger.Fatal(err)
		return "", err
	}
}

// guess the app root path
func guessAppRootPath() (string, error) {
	folder, err := checkAppRoot(osext.ExecutableFolder())
	if err != nil {
		return checkAppRoot(os.Getwd())
	}
	return folder, nil
}

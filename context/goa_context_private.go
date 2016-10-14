package context

import (
	"errors"
	"github.com/arthurlee/goa/file"
	"github.com/kardianos/osext"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func init() {
	//log.Println("GoaContext init")

	folder, err := guessAppRootPath()
	if err != nil {
		log.Fatal("Goa: ", err)
		os.Exit(1)
	}

	Instance.setAppRoot(folder)

	_, err = Instance.loadConfig()
	if err != nil {
		log.Fatal("Goa: ", err)
		os.Exit(1)
	}

}

func (me *GoaContext) setAppRoot(appRoot string) {
	me.AppRootPath = appRoot
	me.AppConfigFilePath = path.Join(appRoot, "conf", "goa.yaml")
}

func (me *GoaContext) loadConfig() (bool, error) {
	yamlContent, err := ioutil.ReadFile(me.AppConfigFilePath)
	if err != nil {
		return false, err
	}

	err = yaml.Unmarshal(yamlContent, &me.Config)
	if err != nil {
		return false, err
	}

	return true, nil
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
		log.Fatal(err)
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

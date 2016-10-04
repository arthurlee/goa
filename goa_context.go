package goa

import (
	"errors"
	"github.com/arthurlee/goa/file"
	"github.com/kardianos/osext"
	"log"
	"os"
	"path"
)

func init() {
	log.Println("goa_context init")

	folder, err := guessAppRootPath()
	if err != nil {
		log.Fatal("Goa: ", err)
		os.Exit(1)
	}

	GoaCtx.setAppRoot(folder)
}

type GoaContext struct {
	AppRootPath       string
	AppConfigFilePath string
}

var GoaCtx GoaContext

func (goaContext *GoaContext) setAppRoot(appRoot string) {
	goaContext.AppRootPath = appRoot
	goaContext.AppConfigFilePath = path.Join(appRoot, "conf", "goa.yaml")
}

// check if goa.yaml exists
func isGoaApp(folder string) bool {
	cfgFilePath := path.Join(folder, "conf", "goa.yaml")
	return fileutil.IsFile(cfgFilePath)
}

// guess the app root path
func guessAppRootPath() (string, error) {
	// exec directory
	folder, err := osext.ExecutableFolder()
	if err == nil {
		if isGoaApp(folder) {
			return folder, nil
		}

		// working directory
		folder, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
			return "", err
		}

		if isGoaApp(folder) {
			return folder, nil
		}

		return "", errors.New("Cannot detect the goa app!")
	} else {
		log.Fatal(err)
		return "", err
	}
}

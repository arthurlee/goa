package goa

import (
	"github.com/arthurlee/goa/database"
	"github.com/arthurlee/goa/instance"
	"github.com/arthurlee/goa/logger"
	"github.com/arthurlee/goa/route"
	"github.com/arthurlee/goa/server"
	"net/http"
)

const GOA_VERSION = "0.0.2"
const GOA_RELEASE_DATE = "2017-02-09"

func showInfo() {
	logger.Info("App root: %s", instance.Instance.AppRootPath)
	logger.Info("Goa %s (%s) service starting", GOA_VERSION, GOA_RELEASE_DATE)
	// It is important to load the database package
	logger.Info("Goa Database version %s", database.GoaDatabaseVersion)
}

// TODO: gracefully exit
func serveHttps(https_addr string, cert_file string, key_file string) {
	logger.Info("https address: %s", https_addr)
	logger.Info("CertFile: %s", cert_file)
	logger.Info("KeyFile: %s", key_file)

	err := http.ListenAndServeTLS(https_addr, cert_file, key_file, route.SrvHandler)
	if err != nil {
		logger.FatalError(err)
	}
	logger.Info("HttpsGoa exit!")
}

func Serve() {
	defer logger.Close()

	showInfo()

	route.Register()

	addr := instance.Instance.Config.Server.Address
	logger.Info("Goa start at address %s", addr)

	// HTTPS support
	https_addr := instance.Instance.Config.Server.HttpsAddress
	if len(https_addr) > 0 {
		cert_file := instance.Instance.Config.Server.CertFile
		key_file := instance.Instance.Config.Server.KeyFile

		go serveHttps(https_addr, cert_file, key_file)
	}

	err := server.HttpListenAndServe(addr, route.SrvHandler)
	if err != nil {
		logger.FatalError(err)
	}
	logger.Info("Goa exit!")
}

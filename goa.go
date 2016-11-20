package goa

import (
	"github.com/arthurlee/goa/database"
	"github.com/arthurlee/goa/instance"
	"github.com/arthurlee/goa/logger"
	"github.com/arthurlee/goa/route"
	"github.com/arthurlee/goa/server"
)

const GOA_VERSION = "0.0.1"
const GOA_RELEASE_DATE = "2016-10-08"

func showInfo() {
	logger.Info("App root: %s", instance.Instance.AppRootPath)
	logger.Info("Goa %s (%s) service starting", GOA_VERSION, GOA_RELEASE_DATE)
	logger.Info("Goa Database version %s", database.GoaDatabaseVersion)
}

func Serve() {
	defer logger.Close()

	showInfo()

	addr := instance.Instance.Config.Server.Address
	logger.Info("Goa start at address %s", addr)
	err := server.HttpListenAndServe(addr, route.SrvHandler)
	if err != nil {
		logger.FatalError(err)
	}
	logger.Info("Goa exit!")
}

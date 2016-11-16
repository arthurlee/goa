package goa

import (
	"github.com/arthurlee/goa/database"
	"github.com/arthurlee/goa/instance"
	"github.com/arthurlee/goa/route"
	"github.com/donnie4w/go-logger/logger"
	"net/http"
)

const GOA_VERSION = "0.0.1"
const GOA_RELEASE_DATE = "2016-10-08"

func showInfo() {
	logger.Info("App root:", instance.Instance.AppRootPath)
	logger.Info("Goa %s (%s) service starting\n", GOA_VERSION, GOA_RELEASE_DATE)

	logger.Info("Goa Database version ", database.GoaDatabaseVersion)
}

func Serve() {
	showInfo()

	addr := instance.Instance.Config.Server.Address
	logger.Info("Goa start at address", addr)
	err := http.ListenAndServe(addr, route.SrvHandler)
	if err != nil {
		logger.Fatal(err)
	}
}

package goa

import (
	"github.com/arthurlee/goa/context"
	"github.com/arthurlee/goa/database"
	"github.com/arthurlee/goa/route"
	"log"
	"net/http"
)

const GOA_VERSION = "0.0.1"
const GOA_RELEASE_DATE = "2016-10-08"

func showInfo() {
	log.Println("App root:", context.Instance.AppRootPath)
	log.Printf("Goa %s (%s) service starting\n", GOA_VERSION, GOA_RELEASE_DATE)

	log.Println("Goa Database version ", database.GoaDatabaseVersion)
}

func Serve() {
	showInfo()
	//database.Init()

	addr := context.Instance.Config.Server.Address
	log.Println("Goa start at address", addr)
	err := http.ListenAndServe(addr, route.SrvHandler)
	if err != nil {
		log.Fatal(err)
	}
}

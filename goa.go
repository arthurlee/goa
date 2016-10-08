package goa

import (
	"github.com/arthurlee/goa/context"
	"log"
	"net/http"
)

const GOA_VERSION = "0.0.1"
const GOA_RELEASE_DATE = "2016-10-08"

func Serve() {
	log.Printf("Goa %s (%s) service starting at %q\n", GOA_VERSION, GOA_RELEASE_DATE, context.Instance.AppRootPath)

	// setup the api server

	addr := context.Instance.Config.Server.Address
	log.Println("Goa start at address", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

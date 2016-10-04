package goa

import (
	"log"
)

func init() {
	log.Println("Goa init")
}

func Serve() {
	log.Println("Goa hello world!")
	log.Println(GoaCtx.AppRootPath)
}

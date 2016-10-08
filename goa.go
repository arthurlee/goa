package goa

import (
	"encoding/json"
	"github.com/arthurlee/goa/context"
	"log"
	"net/http"
)

const GOA_VERSION = "0.0.1"
const GOA_RELEASE_DATE = "2016-10-08"

func init() {
	log.Println("Goa init")
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	person := struct {
		Name string `json:"name", db:"pk"`
		Age  int    `json:"age"`
	}{"Tom", 12}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func Serve() {
	log.Println("Goa %s %s service starting", GOA_VERSION, GOA_RELEASE_DATE)
	log.Println(context.Instance.AppRootPath)

	// setup the api server

	http.HandleFunc("/", handleHello)

	addr := context.Instance.Config.Server.Address
	log.Println("Goa start at address", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

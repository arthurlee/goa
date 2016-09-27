// goa.go
package main

import (
	"encoding/json"
	//	"fmt"
	"log"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	person := struct {
		Name string
		Age  int
	}{"Tom", 12}

	json.NewEncoder(w).Encode(person)
}

func main() {
	http.HandleFunc("/", handleHello)

	port := ":5400"
	log.Println("Goa start at port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

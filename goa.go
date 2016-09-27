// goa.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "[Goa] Hello, world!")
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

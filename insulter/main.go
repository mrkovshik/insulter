// main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	mux:= http.NewServeMux()
	mux.HandleFunc("/name/", Insult)
	mux.HandleFunc("/age/", deathClocker)

	log.Fatal(http.ListenAndServe(":5000", mux))
}

package main

import (
	"log"
	"net/http"

	"github.com/MadaraUchiha-314/cricket/svc/cricketers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cricketers", cricketers.GetCricketers)
	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}
	error := server.ListenAndServe()
	if error != nil {
		log.Fatalf("error while listening to requests: %s", error)
	}
}

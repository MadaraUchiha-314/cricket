package main

import (
	"log"
	"net/http"

	"github.com/MadaraUchiha-314/cricket/api/cricketers"
	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Route("/cricketers", func(router chi.Router) {
		router.Get("/", cricketers.GetCricketers)
	})
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	error := server.ListenAndServe()
	if error != nil {
		log.Fatalf("error while listening to requests: %s", error)
	}
}

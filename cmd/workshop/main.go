package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"workshop/internal/handler"
)

func main() {
	h := handler.NewHandler()
	r := chi.NewRouter()
	r.Get("/", h.Hello)

	log.Printf("starting server")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("shutting server down")
}

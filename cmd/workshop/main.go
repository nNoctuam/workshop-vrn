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

	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}

package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/http"
	"workshop/internal/config"

	"github.com/go-chi/chi"

	"workshop/internal/handler"
)

func main() {

	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler()
	r := chi.NewRouter()
	r.Get("/", h.Hello)

	path := cfg.Host + ":" + cfg.Port
	log.Printf("starting server as " + path)
	err = http.ListenAndServe(path, r)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("shutting server down")
}

package main

import (
	"log"
	"net/http"
	"workshop/internal/api/jokes"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"

	"workshop/internal/config"
	"workshop/internal/handler"
)

func main() {

	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	apiClient := jokes.NewJokeClient(cfg.JokeURL)

	h := handler.NewHandler(apiClient)
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

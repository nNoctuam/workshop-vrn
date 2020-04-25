package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"workshop/internal/api/jokes"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"

	"workshop/internal/config"
	"workshop/internal/handler"
)

func main() {

	ctx := context.Background()

	t := time.NewTicker(3 * time.Second)

	select {
	case <-ctx.Done():
		os.Exit(1)
	case <-t.C:
		fmt.Print("tick")
	}

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

	srv := &http.Server{
		Addr:    path,
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	done := make(chan error, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		ctx, _ := context.WithTimeout(context.Background(), time.Minute)
		done <- srv.Shutdown(ctx)
	}()

	log.Printf("starting server as " + path)
	_ = srv.ListenAndServe()

	err = <-done

	log.Printf("shutting server down with %v", err)
}

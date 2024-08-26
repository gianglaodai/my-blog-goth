package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"my-blog-goth/handlers"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()
	router.Handle("/*", public())
	router.Get("/", handlers.HandleFoo)
	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	err := http.ListenAndServe(listenAddr, router)
	if err != nil {
		log.Fatal(err)
	}
}

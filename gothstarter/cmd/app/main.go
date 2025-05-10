package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()
	// test the router
	router.Get("/foo", handleFoo)
	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listeAddr", listenAddr)
	http.ListenAndServe(listenAddr, router)

}

func handleFoo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("foo"))

}

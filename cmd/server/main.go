package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"MeGrugMeLikeSimple/internal/handlers"
)

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(chimw.Logger)
	r.Use(chimw.Recoverer)
	r.Use(chimw.RealIP)

	// Static files
	workDir, _ := os.Getwd()
	fileServer := http.FileServer(http.Dir(workDir + "/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Routes
	home := handlers.NewHomeHandler()
	r.Get("/", home.ServeHTTP)
	r.Get("/api/hello", home.Hello)
	r.Post("/api/counter", home.Counter)

	// Start server
	addr := ":8080"
	log.Printf("Listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}

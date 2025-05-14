package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/emails", func(r chi.Router) {
		r.Get("/", notImplementedHandler)
		r.Post("/", notImplementedHandler)
	})

	err := http.ListenAndServe(":3333", r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Route not implemented yet"))
}

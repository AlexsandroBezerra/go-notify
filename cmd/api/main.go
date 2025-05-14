package main

import (
	"AlexsandroBezerra/go-notify/internal/api/router"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	ctx := context.Background()

	postgresConnection, err := pgx.Connect(ctx, "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	emailRouter := router.NewEmailRouter(postgresConnection)
	emailRouter.RegisterRoutes(r)

	err = http.ListenAndServe(":3333", r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

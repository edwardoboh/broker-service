package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.AllowContentType("application/json", "text/xml", "text/html"))
	mux.Use(middleware.CleanPath)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://*", "https://*"},
		MaxAge: 300,
		AllowCredentials: false,
		AllowedMethods: []string{"Post", "Get", "Put", "Options", "Patch"},
		ExposedHeaders: []string{"Link"},
		AllowedHeaders: []string{"Content-Type", "Accept", "x-XSRF-Token", "Authorization"},
	}))

	// The Route definations goes here
	mux.Get("/", welcome)

	return mux
}
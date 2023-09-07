package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	Router *mux.Router
}

func NewApp() *App {
	r := mux.NewRouter().
		PathPrefix(os.Getenv("SERVICE_PREFIX")).
		Subrouter().
		StrictSlash(true)

	return &App{Router: r}
}

func (a *App) Listen() error {
	log.Println("Listening on port *:8000")

	srv := &http.Server{
		Handler:      a.Router,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

func (a *App) SetupRoutes() *App {
	router := a.Router

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//
	})

	return a
}

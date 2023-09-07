package main

import (
	"github.com/esportsclub/entity-service-golang/handlers"
	"github.com/esportsclub/entity-service-golang/services"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
	"time"
)

type App struct {
	Router  *mux.Router
	Handler *handlers.Handler
}

func NewApp(db *mongo.Client) *App {
	r := mux.NewRouter().
		PathPrefix(os.Getenv("SERVICE_PREFIX")).
		Subrouter().
		StrictSlash(true)

	service := services.Service{DB: db}
	handler := handlers.Handler{
		Service: &service,
	}

	return &App{Router: r, Handler: &handler}
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
	handler := a.Handler

	router.HandleFunc("/", handler.Home)

	return a
}

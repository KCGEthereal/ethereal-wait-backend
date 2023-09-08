package main

import (
	"github.com/esportsclub/entity-service-golang/database"
	"github.com/esportsclub/entity-service-golang/handlers"
	"github.com/esportsclub/entity-service-golang/services"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
	"time"
)

// App contains all the references that might be used throughout the application.
// We initiate structs and make connections and pass it along to this struct.
// Most of the init happens within NewApp.
//
// Golang has no concept of Dependency injection because those are considered "magical"
// thus it's just plain old passing of values around. You should ideally be able to
// just inspect the references and understand the whole codebase
type App struct {
	router  *mux.Router
	handler *handlers.Handler
}

// NewApp is responsible to take mongodb & redis connection and then also initiate
// other required dependencies for the app and then pass it around via App.
func NewApp(db *mongo.Client, redis *database.Redis) *App {
	r := mux.NewRouter().
		PathPrefix(os.Getenv("SERVICE_PREFIX")).
		Subrouter().
		StrictSlash(true)

	service := services.Service{DB: db, Redis: redis}
	handler := handlers.Handler{
		Service: &service,
	}

	return &App{router: r, handler: &handler}
}

// Listen opens a listener at your desired port and starts accepting requests.
// Always set a WriteTimeout and ReadTimeout inside http.Server
func (a *App) Listen() error {
	log.Println("Listening on port *:8000")

	srv := &http.Server{
		Handler:      a.router,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}

// SetupRoutes houses all the routes used by this application.
func (a *App) SetupRoutes() *App {
	router := a.router
	handler := a.handler

	router.HandleFunc("/", handler.Home)

	return a
}

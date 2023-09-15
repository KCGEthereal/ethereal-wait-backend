// Package handlers contains all the handlers that this project will use.
// Functions receiving Handler type does not need to have handler prefix
// or suffix in its name.
package handlers

import "github.com/KCGEthereal/ethereal-wait-backend/services"

// Handler struct has to be received by every function that is a handler.
// Functions that receive Handler struct as a pointer will have access to
// services.Service this project has. All this has been bootstrapped
// NewApp function that returns App.
//
// Handler is NEVER supposed to have database access. Handler can only
// interact with the http.ResponseWriter and http.Request
type Handler struct {
	Service *services.Service
}

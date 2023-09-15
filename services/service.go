// Package services contains all the services that this project will use.
// Functions receiving Service type does not need to have service prefix
// or suffix in its name.
package services

import (
	"github.com/KCGEthereal/ethereal-wait-backend/database"
	"gorm.io/gorm"
)

// Service struct has to be received by every function that is a service.
// Functions that receive service struct as a pointer will have access to
// database and repository this project has. All this has been bootstrapped
// NewApp function that returns App.
//
// Service will have access to database unlike the Handler. Service holds the
// whole business logic. It should never interact with HTTP req or responses.
// Anything that service requires from client will be built at handler level
// and passed along to the service.
//
// Service will only return required data and error if any
type Service struct {
	DB    *gorm.DB
	Redis *database.Redis
}

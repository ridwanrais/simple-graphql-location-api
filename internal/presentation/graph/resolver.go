package graph

import (
	"github.com/ridwanrais/simple-graphql-location-api/config/database"
	"github.com/ridwanrais/simple-graphql-location-api/internal/repository/mongodb"
	"github.com/ridwanrais/simple-graphql-location-api/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	LocationUsecase usecase.LocationUsecase
}

func NewResolver(locationUsecase usecase.LocationUsecase) *Resolver {
	return &Resolver{
		LocationUsecase: locationUsecase,
	}
}

func InitResolver() *Resolver {
	db := database.InitMongoDB()

	locationRepository := mongodb.NewLocationRepository(db)
	locationUsecase := usecase.NewLocationUsecase(locationRepository)

	return NewResolver(locationUsecase)
}

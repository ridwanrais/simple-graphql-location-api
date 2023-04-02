package mongodb

import (
	"github.com/ridwanrais/simple-graphql-location-api/internal/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type locationRepo struct {
	db *mongo.Database
}

func NewLocationRepository(db *mongo.Database) repository.LocationRepository {
	return &locationRepo{
		db: db,
	}
}

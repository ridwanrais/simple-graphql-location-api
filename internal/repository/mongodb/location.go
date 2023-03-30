package mongodb

import (
	"context"

	"github.com/ridwanrais/simple-graphql-location-api/internal/domain"
	"github.com/ridwanrais/simple-graphql-location-api/internal/repository"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type locationRepo struct {
	db *mongo.Database
}

func NewLocationReposiotry(db *mongo.Database) repository.LocationRepository {
	return &locationRepo{
		db: db,
	}
}

func (r locationRepo) GetCities(ctx context.Context, fields []string) ([]*domain.City, error) {
	projection := bson.M{}

	for _, v := range fields {
		projection[v] = 1
	}

	var cities []*domain.City
	cursor, err := r.db.Collection("cities").Find(ctx, bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		logrus.WithError(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var city domain.City
		if err := cursor.Decode(&city); err != nil {
			logrus.WithError(err)
			return nil, err
		}
		cities = append(cities, &city)
	}

	if err := cursor.Err(); err != nil {
		logrus.WithError(err)
		return nil, err
	}

	return cities, nil
}

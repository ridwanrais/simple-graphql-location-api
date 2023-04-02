package mongodb

import (
	"context"

	"github.com/ridwanrais/simple-graphql-location-api/internal/domain"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r locationRepo) GetCities(ctx context.Context, fields []string, filter domain.CityFilter) ([]*domain.City, error) {
	projection := bson.M{}

	for _, v := range fields {
		projection[v] = 1
	}

	filterCriteria := bson.M{}

	if filter.Name != nil {
		filterCriteria["name"] = *filter.Name
	}

	if filter.CountryID != nil {
		filterCriteria["country_id"] = *filter.CountryID
	}

	if filter.Admin1Code != nil {
		filterCriteria["admin1_code"] = *filter.Admin1Code
	}

	var cities []*domain.City
	cursor, err := r.db.Collection("cities").Find(ctx, filterCriteria, options.Find().SetProjection(projection))
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var city domain.City
		if err := cursor.Decode(&city); err != nil {
			logrus.Error(err)
			return nil, err
		}
		cities = append(cities, &city)
	}

	if err := cursor.Err(); err != nil {
		logrus.Error(err)
		return nil, err
	}

	return cities, nil
}

package mongodb

import (
	"context"

	"github.com/ridwanrais/simple-graphql-location-api/internal/domain"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r locationRepo) GetCountries(ctx context.Context, fields []string) ([]*domain.Country, error) {
	projection := bson.M{}

	for _, v := range fields {
		projection[v] = 1
	}

	cursor, err := r.db.Collection("countries").Find(ctx, bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var countries []*domain.Country
	// for cursor.Next(ctx) {
	// 	var country domain.Country
	// 	if err := cursor.Decode(&country); err != nil {
	// 		logrus.Error(err)
	// 		return nil, err
	// 	}
	// 	countries = append(countries, &country)
	// }

	// if err := cursor.Err(); err != nil {
	// 	logrus.Error(err)
	// 	return nil, err
	// }

	if err = cursor.All(context.Background(), &countries); err != nil {
		logrus.WithError(err)
		return nil, err
	}

	return countries, nil
}

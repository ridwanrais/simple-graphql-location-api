package mongodb

import (
	"context"

	"github.com/ridwanrais/simple-graphql-location-api/internal/domain"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r locationRepo) GetProvinces(ctx context.Context, fields []string, filter domain.Admin1Filter) ([]*domain.Admin1, error) {
	projection := bson.M{}

	for _, v := range fields {
		projection[v] = 1
	}

	filterCriteria := bson.M{}

	if filter.Name != nil {
		filterCriteria["name"] = *filter.Name
	}

	if filter.CountryCCA2 != nil {
		filterCriteria["country_cca2"] = *filter.CountryCCA2
	}

	if filter.Admin1Code != nil {
		filterCriteria["admin1_code"] = *filter.Admin1Code
	}

	cursor, err := r.db.Collection("admin1").Find(ctx, filterCriteria, options.Find().SetProjection(projection))
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var admin1s []*domain.Admin1
	if err = cursor.All(context.Background(), &admin1s); err != nil {
		logrus.WithError(err)
		return nil, err
	}

	return admin1s, nil
}

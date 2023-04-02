package usecase

import (
	"context"

	"github.com/ridwanrais/simple-graphql-location-api/internal/domain"
	"github.com/ridwanrais/simple-graphql-location-api/internal/repository"
	"github.com/ridwanrais/simple-graphql-location-api/shared/converter"
	"github.com/sirupsen/logrus"
)

type LocationUsecase interface {
	GetCities(ctx context.Context, fields []string, filter domain.CityFilter) ([]*domain.City, error)
	GetCountries(ctx context.Context, fields []string) ([]*domain.Country, error)
	GetProvinces(ctx context.Context, fields []string, filter domain.Admin1Filter) ([]*domain.Admin1, error)
}

type locationUsecase struct {
	locationRepo repository.LocationRepository
}

type Pagination struct {
	First  *int
	After  *string
	Last   *int
	Before *string
}

func NewLocationUsecase(locationRepo repository.LocationRepository) LocationUsecase {
	return &locationUsecase{
		locationRepo: locationRepo,
	}
}

func (u locationUsecase) GetCities(ctx context.Context, fields []string, filter domain.CityFilter) ([]*domain.City, error) {
	for i, v := range fields {
		fields[i] = converter.CamelCaseToSnakeCase(v)
	}

	cities, err := u.locationRepo.GetCities(ctx, fields, filter)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return cities, nil
}

func (u locationUsecase) GetCountries(ctx context.Context, fields []string) ([]*domain.Country, error) {
	for i, v := range fields {
		fields[i] = converter.CamelCaseToSnakeCase(v)
	}

	countries, err := u.locationRepo.GetCountries(ctx, fields)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return countries, nil
}

func (u locationUsecase) GetProvinces(ctx context.Context, fields []string, filter domain.Admin1Filter) ([]*domain.Admin1, error) {
	for i, v := range fields {
		fields[i] = converter.CamelCaseToSnakeCase(v)
	}

	admin1s, err := u.locationRepo.GetProvinces(ctx, fields, filter)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return admin1s, nil
}

// func (u locationUsecase) GetCitiesWithPagination(ctx context.Context, fields []string, pagination Pagination) ([]*domain.City, error) {
// 	cities, err := u.locationRepo.GetCities(ctx, fields)
// 	if err != nil {
// 		logrus.Error(err)
// 		return nil, err
// 	}

// 	return cities, nil
// }

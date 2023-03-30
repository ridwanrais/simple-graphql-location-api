package usecase

import (
	"context"

	"github.com/ridwanrais/simple-graphql-location-api/internal/domain"
	"github.com/ridwanrais/simple-graphql-location-api/internal/repository"
	"github.com/sirupsen/logrus"
)

type LocationUsecase interface {
	GetCities(ctx context.Context, fields []string) ([]*domain.City, error)
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

func (u locationUsecase) GetCities(ctx context.Context, fields []string) ([]*domain.City, error) {
	cities, err := u.locationRepo.GetCities(ctx, fields)
	if err != nil {
		logrus.WithError(err)
		return nil, err
	}

	return cities, nil
}

func (u locationUsecase) GetCitiesWithPagination(ctx context.Context, fields []string, pagination Pagination) ([]*domain.City, error) {
	cities, err := u.locationRepo.GetCities(ctx, fields)
	if err != nil {
		logrus.WithError(err)
		return nil, err
	}

	return cities, nil
}

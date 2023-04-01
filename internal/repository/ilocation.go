package repository

import (
	"context"

	"github.com/ridwanrais/simple-graphql-location-api/internal/domain"
)

type LocationRepository interface {
	GetCities(ctx context.Context, fields []string, filter domain.CityFilter) ([]*domain.City, error)
}

package repository

import (
	"context"

	"github.com/ridwanrais/simple-graphql-location-api/internal/domain"
)

type LocationRepository interface {
	GetCities(ctx context.Context, fields []string) ([]*domain.City, error)
}

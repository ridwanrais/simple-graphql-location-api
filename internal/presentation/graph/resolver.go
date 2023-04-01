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
	HelloUseCase    usecase.HelloUseCase
	LocationUsecase usecase.LocationUsecase
}

// func (r *Resolver) Hello(ctx context.Context) ([]*model.Hello, error) {
// 	return r.HelloUseCase.Execute()
// }

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

// func (r *Resolver) Mutation() graph.MutationResolver {
// 	return &mutationResolver{r}
// }

// func (r *Resolver) Query() graph.QueryResolver {
// 	return &queryResolver{r}
// }

// type mutationResolver struct{ *Resolver }
// type queryResolver struct{ *Resolver }

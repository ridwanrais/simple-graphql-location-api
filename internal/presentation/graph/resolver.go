package graph

import (
	"github.com/ridwanrais/simple-graphql-location-api/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	HelloUseCase usecase.HelloUseCase
}

// func (r *Resolver) Hello(ctx context.Context) ([]*model.Hello, error) {
// 	return r.HelloUseCase.Execute()
// }

// func NewResolver() *Resolver {
// 	return &Resolver{
// 		HelloUseCase: usecase.HelloUseCase{},
// 	}
// }

// func (r *Resolver) Mutation() graph.MutationResolver {
// 	return &mutationResolver{r}
// }

// func (r *Resolver) Query() graph.QueryResolver {
// 	return &queryResolver{r}
// }

// type mutationResolver struct{ *Resolver }
// type queryResolver struct{ *Resolver }

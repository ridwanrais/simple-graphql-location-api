package usecase

import "github.com/ridwanrais/simple-graphql-location-api/internal/presentation/graph/model"

type HelloUseCase struct{}

func (u *HelloUseCase) Execute() ([]*model.Hello, error) {
	return []*model.Hello{
		{Text: "Hello"},
	}, nil
}

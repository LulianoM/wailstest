package example

import (
	"context"
	"errors"

	"{{module_name}}/src/interfaces"
	"{{module_name}}/src/models"
	"{{module_name}}/src/repositories"
)

var (
	ErrExampleWithSameID = errors.New("services/example: errors example with same id")
)

type ExampleService struct {
	exampleRepository interfaces.ExampleRepository
}

func NewExampleService(repos repositories.RepositoryContainer) ExampleService {
	return ExampleService{
		exampleRepository: repos.ExampleRepository,
	}
}

func (rc ExampleService) Create(ctx context.Context, example models.Example) (models.Example, error) {
	result, err := rc.exampleRepository.Create(ctx, example)
	if err != nil {
		return models.Example{}, ErrExampleWithSameID
	}
	return result, nil
}

package example

import (
	"context"
	"errors"

	"{{module_name}}/src/models"
)

var (
	ErrExampleNotFound = errors.New("example/repository: errors example not found")
)

type ExampleRepository struct{}

func NewExampleRepository() ExampleRepository {
	return ExampleRepository{}
}

func (rr ExampleRepository) Create(ctx context.Context, example models.Example) (models.Example, error) {
	newExamples := example

	// if err := database.DBClient.GetFromCtx(ctx).Create(&newExamples).Error; err != nil {
	// 	return example, err
	// }
	return newExamples, nil
}

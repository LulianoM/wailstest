package interfaces

import (
	"context"

	"{{module_name}}/src/models"
)

type ExampleRepository interface {
	Create(ctx context.Context, exemple models.Example) (models.Example, error)
}

type ExampleService interface {
	Create(ctx context.Context, exemple models.Example) (models.Example, error)
}

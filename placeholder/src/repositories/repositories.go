package repositories

import (
	"{{module_name}}/src/interfaces"
	"{{module_name}}/src/repositories/example"
)

type RepositoryContainer struct {
	ExampleRepository interfaces.ExampleRepository
}

func NewContainer() RepositoryContainer {
	return RepositoryContainer{
		ExampleRepository: example.NewExampleRepository(),
	}
}

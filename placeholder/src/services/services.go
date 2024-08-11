package services

import (
	"{{module_name}}/src/interfaces"
	"{{module_name}}/src/repositories"
	"{{module_name}}/src/services/example"
	"{{module_name}}/src/services/health"
)

type ServicesContainer struct {
	HealthService  interfaces.HealthService
	ExampleService interfaces.ExampleService
}

func GetServices(repos repositories.RepositoryContainer) ServicesContainer {
	return ServicesContainer{
		HealthService:  health.NewHealthService(),
		ExampleService: example.NewExampleService(repos),
	}
}

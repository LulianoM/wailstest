package example

import (
	"{{module_name}}/src/http/dispatcher"
	"{{module_name}}/src/interfaces"
	"{{module_name}}/src/services"
)

type ExampleHandler struct {
	exampleService interfaces.ExampleService
}

func NewExampleHandler(services services.ServicesContainer) *ExampleHandler {
	return &ExampleHandler{
		exampleService: services.ExampleService,
	}
}

func (rh *ExampleHandler) Routes(d dispatcher.Dispatcher) {
	groupPrivateAPI := d.Private.Group("/example")
	groupPrivateAPI.Post("", rh.CreateExample)
}

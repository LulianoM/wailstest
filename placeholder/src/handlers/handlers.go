package handlers

import (
	"{{module_name}}/src/handlers/example"
	"{{module_name}}/src/handlers/healthcheck"
	"{{module_name}}/src/http/dispatcher"
	"{{module_name}}/src/services"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Routes(dispatcher.Dispatcher)
}

type HandlerContainer struct {
	Handlers []Handler
	Server   *fiber.App
}

func NewHandlerContainer(services services.ServicesContainer) HandlerContainer {
	return HandlerContainer{
		Handlers: []Handler{
			healthcheck.NewHealthHandler(services),
			example.NewExampleHandler(services),
		},
	}
}

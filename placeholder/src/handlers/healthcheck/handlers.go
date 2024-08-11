package healthcheck

import (
	"{{module_name}}/src/http/dispatcher"
	"{{module_name}}/src/interfaces"
	"{{module_name}}/src/services"
)

type HealthHandler struct {
	healthService interfaces.HealthService
}

func NewHealthHandler(services services.ServicesContainer) *HealthHandler {
	return &HealthHandler{
		healthService: services.HealthService,
	}
}

func (hh *HealthHandler) Routes(d dispatcher.Dispatcher) {
	groupApi := d.Public.Group("/health")
	groupApi.Get("/ping", hh.Ping)
}

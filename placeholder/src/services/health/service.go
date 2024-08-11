package health

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (hs *HealthService) Ping() string {
	return "Pong :)"
}

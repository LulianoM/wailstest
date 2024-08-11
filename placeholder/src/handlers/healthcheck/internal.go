package healthcheck

import "github.com/gofiber/fiber/v2"

func (hh *HealthHandler) Ping(c *fiber.Ctx) error {
	s := hh.healthService.Ping()
	return c.SendString(s)

}

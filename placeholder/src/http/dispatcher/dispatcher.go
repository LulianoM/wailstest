package dispatcher

import "github.com/gofiber/fiber/v2"

type Dispatcher struct {
	Public  fiber.Router
	Private fiber.Router
}

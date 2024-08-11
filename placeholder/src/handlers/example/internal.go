package example

import (
	"net/http"

	"{{module_name}}/src/models"
	"github.com/gofiber/fiber/v2"
)

func (rh *ExampleHandler) CreateExample(c *fiber.Ctx) error {
	payload := models.Example{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	exampleCreated, err := rh.exampleService.Create(c.UserContext(), payload)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(models.Response{
		Data: exampleCreated,
	})
}

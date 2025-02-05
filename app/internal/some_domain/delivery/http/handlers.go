package http

import (
	"example-svc/internal/some_domain/delivery"
	"example-svc/internal/some_domain/usecases/models"
	"example-svc/pkg/reqvalidator"

	"github.com/gofiber/fiber/v2"
)

type Example struct {
	SomeDomain delivery.SomeDomain
}

func NewExampleHandlers(uc delivery.SomeDomain) *Example {
	return &Example{SomeDomain: uc}
}

func (e *Example) CreateExample(c *fiber.Ctx) error {
	// some validation
	var req CreateSomeModel
	if err := reqvalidator.ReadRequest(c, &req); err != nil {
		return err
	}

	// call usecase
	id, err := e.SomeDomain.SomeMethodCreate(c.Context(), models.CreateCommand{
		Param1: req.Param1,
		Param2: req.Param2,
	})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"id": id})
}

func (e *Example) GetExample(c *fiber.Ctx) error {
	// some validation
	var req GetExampleModel
	if err := c.QueryParser(&req); err != nil {
		return err
	}

	// call usecase
	model, err := e.SomeDomain.GetExample(c.Context(), ConvertModelToDTO(req))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(model)
}

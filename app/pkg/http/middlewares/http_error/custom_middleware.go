package http_error

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func CreateErrorHandler(errorMap map[error]int) func(*fiber.Ctx, error) CustomError {
	return func(c *fiber.Ctx, err error) CustomError {
		// Going through error map and trying to find appropriate status
		for knownErr, status := range errorMap {
			if errors.Is(err, knownErr) {
				return NewCustomError(c.Status(status).JSON(
					ErrorResponse{Error: knownErr.Error()},
				))
			}
		}

		// If error doesn't find, return empty error
		return NewEmptyCustomError()
	}
}

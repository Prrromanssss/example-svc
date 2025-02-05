package http

import (
	"example-svc/internal/some_domain/usecases"
	errProcessor "example-svc/pkg/http/middlewares/http_error"

	"github.com/gofiber/fiber/v2"
)

var customHandler = errProcessor.CreateErrorHandler(map[error]int{
	usecases.ErrNoAccess: fiber.StatusUnauthorized,
})

func ErrorHandlerMiddleware(c *fiber.Ctx, err error) errProcessor.CustomError {
	return customHandler(c, err)
}

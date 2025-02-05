package http_error

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CustomError struct {
	Msg           error
	IsHaveMessage bool
}

func NewCustomError(msg error) CustomError {
	return CustomError{
		Msg:           msg,
		IsHaveMessage: true,
	}
}

func NewEmptyCustomError() CustomError {
	return CustomError{
		IsHaveMessage: false,
	}
}

type Handler func(c *fiber.Ctx, err error) CustomError

type ErrorProcessor struct {
	errHandlers []Handler
	logger      log.Logger
}

func NewErrorHandler(logger log.Logger, handlers []Handler) *ErrorProcessor {
	if handlers == nil {
		handlers = make([]Handler, 0)
	}
	return &ErrorProcessor{
		errHandlers: handlers,
		logger:      logger,
	}
}

func (e *ErrorProcessor) OverrideErrorHandler(c *fiber.Ctx, err error) error {
	txID, ok := c.Context().Value("requestID").(string)
	if !ok {
		txID = uuid.New().String()
	}

	for _, handler := range e.errHandlers {
		if msg := handler(c, err); msg.IsHaveMessage {
			// logger.Error()
			e.logger.Println("Server catch error",
				"Err", err.Error(),
				"TxID", txID,
			)
			return msg.Msg
		}
	}
	return e.stdError(c, err)
}

func (e *ErrorProcessor) stdError(c *fiber.Ctx, err error) error {
	txID, ok := c.Context().Value("requestID").(string)
	if !ok {
		txID = uuid.New().String()
	}

	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		code = fiberError.Code
	}

	// logger.Error()
	e.logger.Println("Server catch unknown error",
		"Err", err.Error(),
		"TxID", txID,
	)

	// Return status code with error message
	message := c.Status(code).JSON(&ErrorResponse{
		Error: "Server catch unexpected error",
	})
	// Set Content-Type: text/plain; charset=utf-8
	c.Type("json", "utf-8")
	return message
}

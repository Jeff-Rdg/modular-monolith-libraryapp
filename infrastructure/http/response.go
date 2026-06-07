package httpinfra

import (
	"errors"

	govalidator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

const (
	KindNotFound        = "not_found"
	KindConflict        = "conflict"
	KindInputValidation = "input_validation_error"
	KindBusinessRule    = "business_rule_violation"
	KindUnauthorized    = "unauthorized"
	KindForbidden       = "forbidden"
	KindInternalError   = "internal_error"
)

type errorBody struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
}

type validationErrorBody struct {
	Kind   string            `json:"kind"`
	Errors map[string]string `json:"errors"`
}

func ErrorHandler(c fiber.Ctx, err error) error {
	if _, ok := errors.AsType[govalidator.ValidationErrors](err); ok {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": validationErrorBody{
				Kind:   KindInputValidation,
				Errors: ParseErrors(err),
			},
		})
	}

	code := fiber.StatusInternalServerError
	kind := KindInternalError
	msg := "internal server error"

	if fe, ok := errors.AsType[*fiber.Error](err); ok {
		code = fe.Code
		msg = fe.Message
		kind = resolveKind(code)
	}

	return c.Status(code).JSON(fiber.Map{
		"error": errorBody{Kind: kind, Message: msg},
	})
}

func resolveKind(code int) string {
	switch code {
	case fiber.StatusNotFound:
		return KindNotFound
	case fiber.StatusConflict:
		return KindConflict
	case fiber.StatusUnprocessableEntity:
		return KindBusinessRule
	case fiber.StatusUnauthorized:
		return KindUnauthorized
	case fiber.StatusForbidden:
		return KindForbidden
	default:
		return KindInternalError
	}
}

type ResponseBuilder struct {
	c fiber.Ctx
}

func Response(c fiber.Ctx) *ResponseBuilder {
	return &ResponseBuilder{c: c}
}

func (r *ResponseBuilder) Created(data any) error {
	return r.c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": data})
}

func (r *ResponseBuilder) OK(data any) error {
	return r.c.Status(fiber.StatusOK).JSON(fiber.Map{"data": data})
}

func (r *ResponseBuilder) NoContent(msg string) error {
	return r.c.Status(fiber.StatusOK).JSON(fiber.Map{"message": msg})
}

func (r *ResponseBuilder) WithHeader(key, value string) *ResponseBuilder {
	r.c.Set(key, value)
	return r
}

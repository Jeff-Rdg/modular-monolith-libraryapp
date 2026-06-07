package domain

import (
	"github.com/gofiber/fiber/v3"
)

var (
	NotFoundErr     = fiber.NewError(fiber.StatusNotFound, "resource not found")
	ConflictErr     = fiber.NewError(fiber.StatusConflict, "resource already exists")
	BusinessRuleErr = fiber.NewError(fiber.StatusUnprocessableEntity, "business rule violation")
	UnauthorizedErr = fiber.NewError(fiber.StatusUnauthorized, "unauthorized access")
	ForbiddenErr    = fiber.NewError(fiber.StatusForbidden, "forbidden")
)

func New(base *fiber.Error, msg string) *fiber.Error {
	return fiber.NewError(base.Code, msg)
}

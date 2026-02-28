package owner

import (
	"github.com/gofiber/fiber/v2"
)

type OwnerHandler interface {
	Create(c *fiber.Ctx) error
}
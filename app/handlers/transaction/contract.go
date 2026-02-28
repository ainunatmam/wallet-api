package transaction

import (
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler interface {
	Confirmation(c *fiber.Ctx) error
}
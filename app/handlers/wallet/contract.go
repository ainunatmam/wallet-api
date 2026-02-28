package wallet

import (
	"github.com/gofiber/fiber/v2"
)

type WalletHandler interface {
	Create(c *fiber.Ctx) error
	TopUp(c *fiber.Ctx) error
	Payment(c *fiber.Ctx) error
	Transfer(c *fiber.Ctx) error
	Suspend(c *fiber.Ctx) error
	Detail(c *fiber.Ctx) error
	List(c *fiber.Ctx) error
}
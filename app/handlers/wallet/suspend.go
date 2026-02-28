package wallet

import (
	"wallet-api/app/presentation"

	"github.com/gofiber/fiber/v2"
)



func (h *walletHandler) Suspend(c *fiber.Ctx) error {
	walletId := c.Params("walletId")
	if walletId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ResponseBase{}.Failed(fiber.StatusBadRequest, "Invalid Wallet ID"))
	}
	err := h.walletService.Suspend(c.Context(), walletId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ResponseBase{}.Failed(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ResponseBase{}.Success("Success", nil))
}
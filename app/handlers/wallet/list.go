package wallet

import (
	"strconv"
	"wallet-api/app/presentation"

	"github.com/gofiber/fiber/v2"
)



func (h *walletHandler) List(c *fiber.Ctx) error {
	ownerId := c.Params("ownerId")
	if ownerId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ResponseBase{}.Failed(fiber.StatusBadRequest, "Invalid Owner ID"))
	}
	ownerIdUint, err := strconv.ParseUint(ownerId, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ResponseBase{}.Failed(fiber.StatusBadRequest, "Invalid Owner ID"))
	}
	wallets, err := h.walletService.List(c.Context(), ownerIdUint)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ResponseBase{}.Failed(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ResponseBase{}.Success("Success", wallets))
}
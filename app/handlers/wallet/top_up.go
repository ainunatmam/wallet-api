package wallet

import (
	"wallet-api/app/presentation"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)



func (h *walletHandler) TopUp(c *fiber.Ctx) error {
	var req presentation.TransactionTopUpRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ResponseBase{}.Failed(fiber.StatusBadRequest, err.Error()))
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ResponseBase{}.Failed(fiber.StatusBadRequest, err.Error()))
	}

	walletId := c.Params("walletId")
	if walletId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ResponseBase{}.Failed(fiber.StatusBadRequest, "Invalid Wallet ID"))
	}
	result, err := h.walletService.TopUp(c.Context(), walletId, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ResponseBase{}.Failed(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ResponseBase{}.Success("Success", result))
}
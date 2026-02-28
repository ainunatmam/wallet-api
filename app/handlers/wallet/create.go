package wallet

import (
	"wallet-api/app/presentation"
	"wallet-api/app/services/wallet"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type walletHandler struct {
	walletService wallet.WalletService
}

func NewWalletHandler(walletService wallet.WalletService) WalletHandler {
	return &walletHandler{walletService: walletService}
}

func (h *walletHandler) Create(c *fiber.Ctx) error {
	var req presentation.WalletCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ResponseBase{}.Failed(fiber.StatusBadRequest, err.Error()))
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ResponseBase{}.Failed(fiber.StatusBadRequest, err.Error()))
	}

	wallet, err := h.walletService.Create(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ResponseBase{}.Failed(500, err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(presentation.ResponseBase{}.Success("Success", wallet))
}
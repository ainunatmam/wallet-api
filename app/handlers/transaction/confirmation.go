package transaction

import (
	"wallet-api/app/presentation"
	"wallet-api/app/services/transaction"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type transactionHandler struct {
	transactionService transaction.TransactionService
}

func NewTransactionHandler(transactionService transaction.TransactionService) TransactionHandler {
	return &transactionHandler{transactionService: transactionService}
}

func (h *transactionHandler) Confirmation(c *fiber.Ctx) error {
	var req presentation.TransactionConfirmationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ResponseBase{}.Failed(fiber.StatusBadRequest, err.Error()))
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ResponseBase{}.Failed(fiber.StatusBadRequest, err.Error()))
	}

	err := h.transactionService.Confirmation(c, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ResponseBase{}.Failed(500, err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(presentation.ResponseBase{}.Success("Success", nil))
}
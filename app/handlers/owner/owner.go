package owner

import (
	"wallet-api/app/presentation"
	"wallet-api/app/services/owner"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ownerHandler struct	 {
	ownerService owner.OwnerService
}

func NewOwnerHandler(ownerService owner.OwnerService) OwnerHandler {
	return &ownerHandler{
		ownerService: ownerService,
	}
}

func (h *ownerHandler) Create(c *fiber.Ctx) error {
	var req presentation.OwnerCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ResponseBase{}.Failed(fiber.StatusBadRequest, err.Error()))
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(presentation.ResponseBase{}.Failed(fiber.StatusBadRequest, err.Error()))
	}
	
	err := h.ownerService.Create(c, &req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(presentation.ResponseBase{}.Failed(fiber.StatusInternalServerError, err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(presentation.ResponseBase{}.Success("Success", nil))
}
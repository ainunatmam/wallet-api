package owner

import (
	"wallet-api/app/presentation"
	"wallet-api/app/repositories"

	"github.com/gofiber/fiber/v2"
)

type OwnerService interface {
	Create(c *fiber.Ctx, req *presentation.OwnerCreateRequest) error
}

type ownerService struct {
	ownerRepo repositories.OwnerRepository
}

func NewOwnerService(ownerRepo repositories.OwnerRepository) OwnerService {
	return &ownerService{
		ownerRepo: ownerRepo,
	}
}

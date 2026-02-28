package owner

import (
	"wallet-api/app/presentation"

	"github.com/gofiber/fiber/v2"
)

func (s *ownerService) Create(c *fiber.Ctx, req *presentation.OwnerCreateRequest) error {
	err := s.ownerRepo.Create(c.Context(), req.Name)
	if err != nil {
		return err
	}
	return nil
}
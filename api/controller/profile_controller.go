package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/domain"
	"net/http"
)

type ProfileController struct {
	profileUsecase domain.ProfileUsecase
}

func NewProfileController(usecase domain.ProfileUsecase) *ProfileController {
	return &ProfileController{
		profileUsecase: usecase,
	}
}

func (pc *ProfileController) Fetch(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := c.Locals("id").(domain.ID)

	profile, err := pc.profileUsecase.GetProfileByID(ctx, &userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(profile)
}

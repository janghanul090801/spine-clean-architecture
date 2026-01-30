package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/api/controller"
)

func NewSignupRouter(group fiber.Router, controller *controller.SignupController) {
	group.Post("/", controller.Signup)
}

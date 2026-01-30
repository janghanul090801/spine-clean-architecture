package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/api/controller"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/api/middleware"
)

func NewTaskRouter(group fiber.Router, controller *controller.TaskController) {

	// protected
	protected := group.Group("protected")
	protected.Use(middleware.JwtMiddleware)
	protected.Get("/", controller.Fetch)
	protected.Post("/", controller.Create)
}

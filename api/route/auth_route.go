package route

import (
	"github.com/NARUBROWN/spine"
	"github.com/janghanul090801/spine-clean-architecture/api/controller"
)

func NewAuthRouter(app spine.App) {
	app.Route("POST", "/login", (*controller.AuthController).Login)
	app.Route("GET", "/refresh_token", (*controller.AuthController).RefreshToken)
	app.Route("POST", "/signup", (*controller.AuthController).Signup)
}

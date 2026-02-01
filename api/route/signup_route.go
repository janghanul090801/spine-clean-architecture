package route

import (
	"github.com/NARUBROWN/spine"
	"github.com/janghanul090801/spine-clean-architecture/api/controller"
)

func NewSignupRouter(app spine.App) {
	app.Route("POST", "/signup", (*controller.SignupController).Signup)
}

package route

import (
	"github.com/NARUBROWN/spine"
	"github.com/janghanul090801/spine-clean-architecture/api/controller"
)

func NewLoginRouter(app spine.App) {
	app.Route("POST", "/login", (*controller.LoginController).Login)
}

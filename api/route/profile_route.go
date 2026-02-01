package route

import (
	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/pkg/route"
	"github.com/janghanul090801/spine-clean-architecture/api/controller"
	"github.com/janghanul090801/spine-clean-architecture/interceptor"
)

func NewProfileRouter(app spine.App) {
	app.Route("GET", "/profile", (*controller.ProfileController).Fetch, route.WithInterceptors(&interceptor.AuthInterceptor{}))
}

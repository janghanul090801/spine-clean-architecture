package route

import (
	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/pkg/route"
	"github.com/janghanul090801/spine-clean-architecture/api/controller"
	"github.com/janghanul090801/spine-clean-architecture/interceptor"
)

func NewRefreshTokenRouter(app spine.App) {
	app.Route("GET", "/refresh_token", (*controller.RefreshTokenController).RefreshToken, route.WithInterceptors(&interceptor.AuthInterceptor{}))
}

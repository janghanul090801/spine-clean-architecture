package route

import (
	"github.com/NARUBROWN/spine"
	"github.com/NARUBROWN/spine/pkg/route"
	"github.com/janghanul090801/spine-clean-architecture/api/controller"
	"github.com/janghanul090801/spine-clean-architecture/interceptor"
)

func NewTaskRouter(app spine.App) {
	app.Route("GET", "/task", (*controller.TaskController).Fetch, route.WithInterceptors(&interceptor.AuthInterceptor{}))
	app.Route("POST", "/task", (*controller.TaskController).Create, route.WithInterceptors(&interceptor.AuthInterceptor{}))
}

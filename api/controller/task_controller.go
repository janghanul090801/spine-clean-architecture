package controller

import (
	"context"
	"github.com/NARUBROWN/spine/pkg/httperr"
	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/NARUBROWN/spine/pkg/spine"
	"github.com/janghanul090801/spine-clean-architecture/domain"
	"net/http"
)

type TaskController struct {
	taskUsecase domain.TaskUsecase
}

func NewTaskController(usecase domain.TaskUsecase) *TaskController {
	return &TaskController{
		taskUsecase: usecase,
	}
}

func (tc *TaskController) Create(ctx context.Context, task *domain.Task, spineCtx spine.Ctx) error {

	v, ok := spineCtx.Get("id")
	if !ok {
		return httperr.Unauthorized("unauthorized")
	}

	userID := v.(domain.ID)

	task.UserID = userID

	err := tc.taskUsecase.Create(ctx, task)
	if err != nil {
		return &httperr.HTTPError{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Cause:   err,
		}
	}

	return nil
}

func (tc *TaskController) Fetch(ctx context.Context, spineCtx spine.Ctx) httpx.Response[[]domain.Task] {
	v, ok := spineCtx.Get("id")
	if !ok {
		return httpx.Response[[]domain.Task]{
			Options: httpx.ResponseOptions{
				Status: http.StatusUnauthorized, // unauthorized
			},
		}
	}

	userID := v.(domain.ID)

	tasks, err := tc.taskUsecase.FetchByUserID(ctx, &userID)
	if err != nil {
		return httpx.Response[[]domain.Task]{
			Options: httpx.ResponseOptions{
				Status: http.StatusInternalServerError, // err.Error()
			},
		}
	}

	taskValues := make([]domain.Task, len(tasks))
	for i, task := range tasks {
		taskValues[i] = domain.Task{
			ID:        task.ID,
			Title:     task.Title,
			UserID:    task.UserID,
			CreatedAt: task.CreatedAt,
		}
	}

	return httpx.Response[[]domain.Task]{
		Body: taskValues,
	}
}

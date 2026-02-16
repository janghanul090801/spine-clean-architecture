package controller

import (
	"context"
	"errors"
	"net/http"

	"github.com/NARUBROWN/spine/pkg/httperr"
	"github.com/NARUBROWN/spine/pkg/httpx"
	"github.com/NARUBROWN/spine/pkg/spine"
	"github.com/janghanul090801/spine-clean-architecture/domain"
)

type TaskController struct {
	taskUseCase domain.TaskUseCase
}

func NewTaskController(usecase domain.TaskUseCase) *TaskController {
	return &TaskController{
		taskUseCase: usecase,
	}
}

func (tc *TaskController) Create(ctx context.Context, task *domain.Task, spineCtx spine.Ctx) error {
	var errInfo domain.Error

	v, ok := spineCtx.Get("id")
	if !ok {
		return httperr.Unauthorized("unauthorized")
	}

	userID := v.(*domain.ID)

	_, err := tc.taskUseCase.Create(ctx, task, userID)
	if err != nil {
		if ok := errors.As(err, &errInfo); ok {
			return &httperr.HTTPError{
				Status:  errInfo.StatusCode,
				Message: err.Error(),
				Cause:   err,
			}
		}
	}

	return nil
}

func (tc *TaskController) Fetch(ctx context.Context, spineCtx spine.Ctx) httpx.Response[[]domain.Task] {
	var errInfo domain.Error

	v, ok := spineCtx.Get("id")
	if !ok {
		return httpx.Response[[]domain.Task]{
			Options: httpx.ResponseOptions{
				Status: http.StatusUnauthorized, // unauthorized
			},
		}
	}

	userID := v.(*domain.ID)

	tasks, err := tc.taskUseCase.FetchByUserID(ctx, userID)
	if err != nil {
		if ok := errors.As(err, &errInfo); ok {
			return httpx.Response[[]domain.Task]{
				Options: httpx.ResponseOptions{
					Status: errInfo.StatusCode, // err.Error()
				},
			}
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

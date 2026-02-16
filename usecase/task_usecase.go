package usecase

import (
	"context"
	"time"

	"github.com/janghanul090801/spine-clean-architecture/domain"
)

type taskUseCase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUseCase(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUseCase {
	return &taskUseCase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (tu *taskUseCase) Create(c context.Context, task *domain.Task, userID *domain.ID) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	task.UserID = *userID

	t, err := tu.taskRepository.Create(ctx, task)
	if err != nil {
		return nil, domain.NewInternalServerError(err)
	}

	return t, nil
}

func (tu *taskUseCase) FetchByUserID(c context.Context, userID *domain.ID) ([]*domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	tasks, err := tu.taskRepository.FetchByUserID(ctx, userID)
	if err != nil {
		return nil, domain.NewInternalServerError(err)
	}

	return tasks, nil
}

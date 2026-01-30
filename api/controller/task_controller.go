package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/bootstrap"
	"github.com/janghanul090801/go-backend-clean-architecture-fiber/domain"
	"net/http"
)

type TaskController struct {
	taskUsecase domain.TaskUsecase
}

func NewTaskController(usecase domain.TaskUsecase, env *bootstrap.Env) *TaskController {
	return &TaskController{
		taskUsecase: usecase,
	}
}

func (tc *TaskController) Create(c *fiber.Ctx) error {
	ctx := c.Context()
	var task domain.Task

	err := c.BodyParser(&task)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	userID := c.Locals("id").(domain.ID)

	task.UserID = userID

	err = tc.taskUsecase.Create(ctx, &task)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(domain.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (tc *TaskController) Fetch(c *fiber.Ctx) error {
	ctx := c.Context()
	userID := c.Locals("id").(domain.ID)

	tasks, err := tc.taskUsecase.FetchByUserID(ctx, &userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(domain.ErrorResponse{Message: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(tasks)
}

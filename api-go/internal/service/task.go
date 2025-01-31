package service

import (
	"todo-app/internal/config"
	"todo-app/internal/model"
	"todo-app/internal/repository"

	"github.com/labstack/echo/v4"
)

func NewTaskService(env *config.Env, repository *repository.Repository) *TaskService {
	return &TaskService{
		env: env,
		taskRepo: repository.TaskRepo,
	}
}

type TaskService struct {
	env *config.Env
	taskRepo *repository.TaskRepo
}

func (ts *TaskService) Create(ctx echo.Context, data TaskCreateInput) error {
	return ts.taskRepo.Create(ctx, repository.TaskCreateInput{
		UserId: data.UserId,
		Title: data.Title,
	})
}

func (ts *TaskService) GetLists(ctx echo.Context, userId int64, limit int, offset int) ([]model.Task, error) {
	return ts.taskRepo.GetTasks(ctx, userId, limit, offset)
}

func (ts *TaskService) GetTotal(ctx echo.Context, userId int64) (int64, error) {
	return ts.taskRepo.GetTasksTotal(ctx, userId)
}


func (ts *TaskService) CompleteTask(ctx echo.Context, userId int64, taskId int64, ) error {
	return ts.taskRepo.CompleteTask(ctx, userId, taskId)
}
package service

import (
	"todo-app/internal/config"
	"todo-app/internal/repository"
)

type Service struct {
	UserService *UserService
	TaskService *TaskService
}

func NewService(env *config.Env, repository *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(env, repository),
		TaskService: NewTaskService(env, repository),
	}
}
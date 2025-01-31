package repository

import (
	"database/sql"
	"todo-app/internal/config"
)

type Repository struct {
	UserRepo *UserRepo
	TaskRepo *TaskRepo
}

func NewRepo(env *config.Env, db *sql.DB) *Repository {
	return &Repository{
		UserRepo: NewUserRepo(env, db),
		TaskRepo: NewTaskRepo(env, db),
	}
}
package repository

import (
	"database/sql"
	"time"
	"todo-app/internal/config"
	"todo-app/internal/constant"
	"todo-app/internal/libs"
	"todo-app/internal/model"

	"github.com/labstack/echo/v4"
)

func NewTaskRepo(env *config.Env, db *sql.DB) *TaskRepo {
	return &TaskRepo{
		env: env,
		db: db,
	}
}

type TaskRepo struct {
	env *config.Env
	db *sql.DB
}

func (tr *TaskRepo) Create(ctx echo.Context, data TaskCreateInput) error {
	now := time.Now().Format(time.DateTime)

	query := `
	INSERT INTO tasks (user_id, task_title, task_created_at, task_updated_at)
	VALUES ($1, $2, $3, $4);
	`

	_, err := tr.db.Exec(query, data.UserId, data.Title, now, now)
	if err != nil {
		return libs.DefaultInternalServerError(ctx, err)
	}

	return nil
}

func (tr *TaskRepo) CompleteTask(ctx echo.Context, userId int64, taskId int64) error {
	now := time.Now().Format(time.DateTime)

	query := `
	UPDATE tasks
	SET
		task_status = $1,
		task_updated_at = $2
	WHERE
		task_id = $3 AND 
		user_id = $4; 
	`

	_, err := tr.db.Exec(query, constant.TASK_STATUS_COMPLETED, now, taskId, userId)
	if err != nil {
		return libs.DefaultInternalServerError(ctx, err)
	}

	return nil
}

func (tr *TaskRepo) GetTasks(ctx echo.Context, userId int64, limit int, offset int) ([]model.Task, error) {
	query := `
	SELECT task_id, user_id, task_title, task_status, task_created_at, task_updated_at
	FROM tasks
	WHERE user_id = $1 AND task_status = $2
	ORDER BY task_id DESC
	LIMIT $3
	OFFSET $4;
	`

	rows, err := tr.db.Query(query, userId, constant.TASK_STATUS_ON_GOING, limit, offset)
	if err != nil {
		return nil, libs.DefaultInternalServerError(ctx, err)
	}

	defer rows.Close()
	tasks := []model.Task{}
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return tasks, libs.DefaultInternalServerError(ctx, err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (tr *TaskRepo) GetTasksTotal(ctx echo.Context, userId int64) (int64, error) {
	query := `
	SELECT COUNT(task_id)
	FROM tasks
	WHERE user_id = $1 AND task_status = $2;
	`

	var totalTask int64
	err := tr.db.QueryRow(query, userId, constant.TASK_STATUS_ON_GOING).Scan(&totalTask)
	if err != nil {
		return 0, libs.DefaultInternalServerError(ctx, err)
	}

	return totalTask, nil
}
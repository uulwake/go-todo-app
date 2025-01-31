package repository

import (
	"database/sql"
	"time"
	"todo-app/internal/config"
	"todo-app/internal/libs"
	"todo-app/internal/model"

	"github.com/labstack/echo/v4"
)

func NewUserRepo(env *config.Env, db *sql.DB) *UserRepo {
	return &UserRepo{
		env: env,
		db: db,
	}
}

type UserRepo struct {
	env *config.Env
	db *sql.DB
}

func (ur *UserRepo) Create(ctx echo.Context, data UserCreateInput) (int64, error) {
	now := time.Now().Format(time.DateTime)

	query := `
	INSERT INTO users (user_name, user_email, user_password, user_created_at, user_updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING user_id;
	`

	var userID int64
	err := ur.db.QueryRow(query, data.Name, data.Email, data.Password, now, now).Scan(&userID)

	if err != nil {
		return 0, libs.DefaultInternalServerError(ctx, err)
	}
	
	return userID, nil
}

func (ur *UserRepo) GetPasswordByEmail(ctx echo.Context, email string) (model.User, error) {
	query := `
	SELECT user_id, user_name, user_email, user_password
	FROM users
	WHERE user_email = $1
	LIMIT 1;
	`

	var user model.User
	err := ur.db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return model.User{}, libs.DefaultInternalServerError(ctx, err)
	}

	return user, nil
}
package service

import (
	"net/http"
	"time"
	"todo-app/internal/config"
	"todo-app/internal/libs"
	"todo-app/internal/model"
	"todo-app/internal/repository"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func NewUserService(env *config.Env, repo *repository.Repository) *UserService {
	return &UserService{
		env: env,
		userRepo: repo.UserRepo,
	}
}

type UserService struct {
	env *config.Env
	userRepo *repository.UserRepo
}

func (us *UserService) CreateJwtToken(ctx echo.Context, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp": time.Now().Add(time.Duration(us.env.JwtExpired * int(time.Hour))).Unix(),
	})

	stringToken, err := token.SignedString([]byte(us.env.JwtSecret))
	if err != nil {
		return "", err
	}

	return stringToken, nil
}

func (us *UserService) Register(ctx echo.Context, data UserRegisterInput) (int64, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(data.Password), us.env.Salt)
	if err != nil {
		return 0, libs.CustomError{
			HTTPCode: http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return us.userRepo.Create(ctx, repository.UserCreateInput{
		Name: data.Name,
		Email: data.Email,
		Password: string(hashedPwd),
	})
}

func (us *UserService) Login(ctx echo.Context, data UserLoginInput) (model.User, error) {
	user, err := us.userRepo.GetPasswordByEmail(ctx, data.Email)

	if err != nil {
		return model.User{}, libs.CustomError{
			HTTPCode: http.StatusBadRequest,
			Message: "Invalid email/password",
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return model.User{}, libs.CustomError{
			HTTPCode: http.StatusBadRequest,
			Message: "Invalid email/password",
		}
	}

	return user, nil
}
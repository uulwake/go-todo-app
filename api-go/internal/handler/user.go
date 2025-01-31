package handler

import (
	"net/http"
	"todo-app/internal/config"
	"todo-app/internal/libs"
	"todo-app/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func InitUserHandler(echoGroup *echo.Group, env *config.Env, validate *validator.Validate,service *service.Service) {
	uh := &UserHandler{
		echoGroup: echoGroup, 
		env: env, 
		validate: validate,
		userService: service.UserService,
	}

	uh.echoGroup.POST("/register", uh.Register)
	uh.echoGroup.POST("/login", uh.Login)
}

type UserHandler struct {
	echoGroup *echo.Group
	env *config.Env
	validate *validator.Validate
	userService *service.UserService
}

func (uh *UserHandler) Register(ctx echo.Context) error {
	var body UserRegisterRequest

	err := PreprocessedRequest(ctx, uh.validate, &body)
	if err != nil {
		return libs.CustomError{
			HTTPCode: http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	userID, err := uh.userService.Register(ctx, service.UserRegisterInput{Name: body.Name, Email: body.Email, Password: body.Password})
	if err != nil {
		return err
	}

	jwtToken, err := uh.userService.CreateJwtToken(ctx, userID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, UserRegisterLoginResponse{
		Data: UserRegisterLoginResponseData{
			UserID: userID,
			UserName: body.Name,
			Token: jwtToken,
		},
	})
}

func (uh *UserHandler) Login(ctx echo.Context) error {
	var body UserLoginRequest

	err := PreprocessedRequest(ctx, uh.validate, &body)
	if err != nil {
		return libs.CustomError{
			HTTPCode: http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	user, err := uh.userService.Login(ctx, service.UserLoginInput{
		Email: body.Email,
		Password: body.Password,
	})
	if err != nil {
		return err
	}


	jwtToken, err := uh.userService.CreateJwtToken(ctx, user.ID)
	if err != nil {
		return err
	}


	return ctx.JSON(http.StatusOK, UserRegisterLoginResponse{
		Data: UserRegisterLoginResponseData{
			UserID: user.ID,
			UserName: user.Name,
			Token: jwtToken,
		},
	})
}
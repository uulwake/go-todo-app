package handler

import (
	"net/http"
	"todo-app/internal/config"
	"todo-app/internal/handler/middleware"
	"todo-app/internal/libs"
	"todo-app/internal/model"
	"todo-app/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func InitTaskHandler(echoGroup *echo.Group, env *config.Env, validate *validator.Validate, service *service.Service) {
	th := &TaskHandler{
		echoGroup: echoGroup,
		env: env,
		validate: validate,
		taskService: service.TaskService,
	}

	th.echoGroup.Use(middleware.AuthenticateJwt(env))
	th.echoGroup.POST("", th.Create)
	th.echoGroup.GET("", th.GetLists)
	th.echoGroup.PATCH("/:task_id", th.CompleteTask)
}

type TaskHandler struct {
	echoGroup *echo.Group
	env *config.Env
	validate *validator.Validate
	taskService *service.TaskService
}

func (th *TaskHandler) Create(ctx echo.Context) error {
	var body TaskCreateBody
	err := PreprocessedRequest(ctx, th.validate, &body)
	if err != nil {
		return err
	}

	userId, err := GetUserIdFromContext(ctx)
	if err != nil {
		return err
	}
	
	err = th.taskService.Create(ctx, service.TaskCreateInput{
		UserId: userId,
		Title: body.Title,
	})

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, struct{}{})
}

func (th *TaskHandler) GetLists(ctx echo.Context) error {
	err := PreprocessedRequest(ctx, th.validate, nil)
	if err != nil {
		return err
	}

	userId, err := GetUserIdFromContext(ctx)
	if err != nil {
		return err
	}

	commonQueryParam, err := ConvertCommonQueryParam(ctx)
	if err != nil {
		return err
	}

	limit := commonQueryParam.PageSize
	offset := (commonQueryParam.PageNumber-1)*limit


	getListsOutputChan := make(chan []model.Task, 1)
	getTotalOutputChan := make(chan int64, 1)
	errorChan := make(chan error, 2)

		

	go func(getListsOutputChan chan<- []model.Task, errorChan chan<- error, ctx echo.Context, userId int64, limit int, offset int) {
		tasks, err := th.taskService.GetLists(ctx, userId, limit, offset)

		getListsOutputChan <- tasks
		errorChan <- err
		
		close(getListsOutputChan)
	}(getListsOutputChan, errorChan, ctx, userId, limit, offset)


	go func(getTotalOutputChan chan<- int64, errorChan chan<- error, ctx echo.Context, userId int64) {
		total, err := th.taskService.GetTotal(ctx, userId)

		getTotalOutputChan <- total
		errorChan <- err

		close(getTotalOutputChan)
	}(getTotalOutputChan, errorChan, ctx, userId)


	tasks := <-getListsOutputChan
	total := <-getTotalOutputChan

	err1 := <-errorChan
	err2 := <-errorChan 
	close(errorChan)

	if err1 != nil {
		return err1
	}

	if err2 != nil {
		return err2
	}


	return ctx.JSON(http.StatusOK, TaskGetListsResponse{
		Data: tasks,
		Page: libs.Pagination{
			Size: limit,
			Number: offset,
			Total: total,
		},
	})	
}

func (th *TaskHandler) CompleteTask(ctx echo.Context) error {
	userId, err := GetUserIdFromContext(ctx)
	if err != nil {
		return err
	}

	taskId, err := GetIdFromPathParam(ctx, "task_id")
	if err != nil {
		return err
	}

	err = th.taskService.CompleteTask(ctx, userId, taskId)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, struct{}{})
}
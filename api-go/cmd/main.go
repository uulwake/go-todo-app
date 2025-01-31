package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"todo-app/internal/config"
	"todo-app/internal/database"
	"todo-app/internal/handler"
	"todo-app/internal/repository"
	"todo-app/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	env, err := config.NewEnv()
	if err != nil {
		log.Fatal("Failed reading environment variables: ", err)
	}

	db, err := database.NewPg(env)
	if err != nil {
		log.Fatal("Failed connect to database", err)
	}

	repository := repository.NewRepo(env, db)
	service := service.NewService(env, repository)

	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(middleware.CORS())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		ErrorMessage: "request timeout",
		Timeout: 30 * time.Second,
	}))
	e.IPExtractor = echo.ExtractIPFromRealIPHeader()
	e.HTTPErrorHandler = handler.CustomHTTPErrorHandler

	e.GET("/hc", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string `json:"status"`}{Status: "ok"})
	})
	
	validate := validator.New()

	v1Group := e.Group("/v1")
	handler.InitUserHandler(v1Group.Group("/users"), env, validate, service)
	handler.InitTaskHandler(v1Group.Group("/tasks"), env, validate, service)

	// run
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", env.Port)))

}
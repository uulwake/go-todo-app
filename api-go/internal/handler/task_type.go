package handler

import (
	"todo-app/internal/libs"
	"todo-app/internal/model"
)

type TaskCreateBody struct {
	Title string `json:"title" validation:"required,min=1"`
}

type TaskGetListsQueryParam struct {
	libs.QueryParam
}

type TaskGetListsResponse struct {
	Data []model.Task `json:"data"`
	Page libs.Pagination `json:"page"`
}

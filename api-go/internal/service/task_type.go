package service

import "todo-app/internal/libs"

type TaskCreateInput struct {
	UserId int64
	Title string
}

type TaskGetListsQueryParam struct {
	libs.QueryParam
}
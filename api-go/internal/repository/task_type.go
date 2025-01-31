package repository

type TaskCreateInput struct {
	UserId int64
	Title string
}

type TaskGetListsQuery struct {
	Limit int
	Offset int
}
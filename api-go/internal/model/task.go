package model

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Task struct {
	ID int64 `json:"task_id,omitempty"`
	UserID int64 `json:"user_id,omitempty"`
	Title string `json:"task_title,omitempty"`
	Description *null.String `json:"task_description,omitempty"`
	Status string `json:"task_status,omitempty"`
	CreatedAt *time.Time `json:"task_created_at,omitempty"`
	UpdatedAt *time.Time `json:"task_updated_at,omitempty"`
}
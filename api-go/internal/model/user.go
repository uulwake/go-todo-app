package model

import "time"

type User struct {
	ID int64 `json:"user_id,omitempty"`
	Name string `json:"user_name,omitempty"`
	Email string `json:"user_email,omitempty"`
	Password string `json:"user_password,omitempty"`
	CreatedAt *time.Time `json:"user_created_at,omitempty"`
	UpdatedAt *time.Time `json:"user_updated_at,omitempty"`
}
package requests

import (
	"time"
)

type UserRequest struct {
	First_name      string `json:"first_name" validate:"required"`
	Last_name       string `json:"last_name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Profile_picture string `json:"profile_picture"`
	Phone_number    string `json:"phone_number" validate:"required"`
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	Role_id         uint
}

type UpdateUserProfileRequest struct {
	First_name      string    `json:"first_name" validate:"required"`
	Last_name       string    `json:"last_name" validate:"required"`
	Username        string    `json:"username" validate:"required"`
	Email           string    `json:"email"  validate:"required"`
	Birthday        time.Time `json:"birthday"`
	Profile_picture string    `json:"profile_picture"`
	Role_id         uint
}

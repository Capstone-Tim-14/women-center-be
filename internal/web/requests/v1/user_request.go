package requests

type UserRequest struct {
	First_name      string `json:"first_name" validate:"required"`
	Last_name       string `json:"last_name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Profile_picture string `json:"profile_picture"`
	Phone_number    string `json:"phone_number" validate:"required"`
	Address         string `json:"address" validate:"required"`
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	Role_id         uint
}

type UpdateUserProfileRequest struct {
	First_name      string `json:"first_name" validate:"required" form:"first_name"`
	Last_name       string `json:"last_name" validate:"required" form:"last_name"`
	Username        string `json:"username" validate:"required" form:"username"`
	Email           string `json:"email"  validate:"required" form:"email"`
	Birthday        string `json:"birthday" form:"birthday"`
	Profile_picture string `json:"profile_picture" form:"profile_picture"`
	Role_id         uint
}

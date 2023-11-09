package requests

type UserRequest struct {
	First_name      string `json:"first_name" validate:"required"`
	Last_name       string `json:"last_name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Profile_picture string `json:"profile_picture"`
	Phone_number    int    `json:"phone_number" validate:"required"`
	Address         string `json:"address" validate:"required"`
	// Status          string `json:"status" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role_id  uint   `json:"role" validate:"required"`
}

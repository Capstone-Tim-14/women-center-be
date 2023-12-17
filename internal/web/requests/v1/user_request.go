package requests

type UserRequest struct {
	First_name      string `json:"first_name" validate:"required" form:"first_name"`
	Last_name       string `json:"last_name" validate:"required" form:"last_name"`
	Email           string `json:"email" validate:"required,email" form:"email"`
	Profile_picture string `json:"profile_picture" form:"profile_picture"`
	Phone_number    string `json:"phone_number" validate:"required" form:"phone_number"`
	Username        string `json:"username" validate:"required" form:"username"`
	Password        string `json:"password" validate:"required,min=6" form:"password"`
	Role_id         uint
}

type UpdateUserProfileRequest struct {
	First_name      string `json:"first_name" validate:"required" form:"first_name"`
	Last_name       string `json:"last_name" validate:"required" form:"last_name"`
	Username        string `json:"username" validate:"required" form:"username"`
	Email           string `json:"email"  validate:"required" form:"email"`
	Birthday        string `json:"birthday" form:"birthday" validate:"required"`
	Profile_picture string `json:"profile_picture" form:"profile_picture"`
	Role_id         uint
}

type GenerateOTPRequest struct {
	Email string `json:"email"`
}

type VerifyOTPRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

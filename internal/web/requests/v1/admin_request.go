package requests

type AdminRequest struct {
	First_name string `json:"first_name" validate:"required"`
	Last_name  string `json:"last_name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	Role_id    uint
}

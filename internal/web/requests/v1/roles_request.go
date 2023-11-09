package requests

type RoleRequest struct {
	Name string `json:"name" validate:"required"`
}

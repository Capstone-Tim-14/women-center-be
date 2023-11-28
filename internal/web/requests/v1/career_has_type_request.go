package requests

type CareerhasTypeRequest struct {
	Name string `json:"name" validate:"required"`
}

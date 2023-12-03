package requests

type CareerhasTypeRequest struct {
	Name string `json:"name" validate:"required"`
}

type CareerhasManyRequest struct {
	Name []string `json:"name" validate:"required"`
}

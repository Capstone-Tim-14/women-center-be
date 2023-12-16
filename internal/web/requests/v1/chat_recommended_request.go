package requests

type ChatRecomendedRequest struct {
	Message string `json:"message" validate:"required"`
}

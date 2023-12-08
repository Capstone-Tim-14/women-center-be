package requests

type CounselorFavotireRequest struct {
	Name string `json:"name" validate:"required"`
}

package requests

type CounselorHasSpecialistRequest struct {
	Name string `json:"name" validate:"required"`
}

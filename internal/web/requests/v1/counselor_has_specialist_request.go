package requests

type CounselorHasSpecialistRequest struct {
	Name string `json:"name" validate:"required"`
}

type CounselorHasManyRequest struct {
	Name []string `json:"name" validate:"required"`
}

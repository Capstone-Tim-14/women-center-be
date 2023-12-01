package requests

type CounselingPackageRequest struct {
	Package_name       string  `json:"package_name" validate:"required" form:"package_name"`
	Description        string  `json:"description" validate:"required" form:"description"`
	Thumbnail          *string `json:"thumbnail" validate:"required" form:"thumbnail"`
	Number_of_sessions uint    `json:"number_of_sessions" validate:"required" form:"number_of_sessions"`
	Price              string  `json:"price" validate:"required" form:"price"`
}

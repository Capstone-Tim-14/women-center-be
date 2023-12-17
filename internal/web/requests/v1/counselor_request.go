package requests

type CounselorRequest struct {
	First_name      string `json:"first_name" validate:"required" form:"first_name"`
	Last_name       string `json:"last_name" validate:"required" form:"last_name"`
	Email           string `json:"email" validate:"required,email" form:"email"`
	Profile_picture string `json:"profile_picture" form:"profile_picture"`
	Description     string `json:"description" validate:"required" form:"description"`
	Username        string `json:"username" validate:"required" form:"username"`
	Password        string `json:"password" validate:"required,min=6" form:"password"`
	Role_id         uint
}

type FilterCounselorsSpecialist struct {
	Specialist []string `json:"specialist" form:"specialist"`
}

type UpdateCounselorProfileRequest struct {
	First_name      string `json:"first_name" validate:"required" form:"first_name"`
	Last_name       string `json:"last_name" validate:"required" form:"last_name"`
	Username        string `json:"username" validate:"required" form:"username"`
	Email           string `json:"email"  validate:"required" form:"email"`
	Password        string `json:"password" validate:"required" form:"password"`
	Profile_picture string `json:"profile_picture" form:"profile_picture"`
	Description     string `json:"description" validate:"required" form:"description"`
	Education       string `json:"education" validate:"required" form:"education"`
	Birthday        string `json:"birthday" form:"birthday" validate:"required"`
	Role_id         uint
}

type UpdateCounselorProfileRequestForMobile struct {
	First_name      string `json:"first_name" validate:"required" form:"first_name"`
	Last_name       string `json:"last_name" validate:"required" form:"last_name"`
	Email           string `json:"email"  validate:"required" form:"email"`
	Profile_picture string `json:"profile_picture" form:"profile_picture"`
	Birthday        string `json:"birthday" form:"birthday" validate:"required"`
	Role_id         uint
}

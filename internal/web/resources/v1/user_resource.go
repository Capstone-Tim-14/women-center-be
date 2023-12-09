package resources

type UserResource struct {
	Id              uint   `json:"id,omitempty"`
	First_name      string `json:"first_name,omitempty"`
	Last_name       string `json:"last_name,omitempty"`
	Email           string `json:"email,omitempty"`
	Username        string `json:"username,omitempty"`
	Profile_picture string `json:"profile_picture,omitempty"`
	Phone_number    string `json:"phone_number,omitempty"`
	Status          string `json:"status,omitempty"`
}

type GetUserProfile struct {
	Id              uint   `json:"id,omitempty"`
	Profile_picture string `json:"profile_picture,omitempty"`
	Username        string `json:"username,omitempty"`
	Full_name       string `json:"full_name,omitempty"`
	Email           string `json:"email,omitempty"`
	Birthday        string `json:"birthday,omitempty"`
}

type UpdateUserProfile struct {
	Id              uint   `json:"id,omitempty"`
	First_name      string `json:"first_name,omitempty"`
	Last_name       string `json:"last_name,omitempty"`
	Username        string `json:"username,omitempty"`
	Email           string `json:"email,omitempty"`
	Birthday        string `json:"birthday,omitempty"`
	Profile_picture string `json:"profile_picture,omitempty"`
}

type UserCounselorFavorite struct {
	Id                uint                `json:"id,omitempty"`
	First_name        string              `json:"first_name,omitempty"`
	Last_name         string              `json:"last_name,omitempty"`
	Username          string              `json:"username,omitempty"`
	CounselorFavorite []CounselorFavorite `json:"counselor_favorite,omitempty"`
}

type CounselorFavorite struct {
	Id              uint   `json:"id,omitempty"`
	First_name      string `json:"first_name,omitempty"`
	Last_name       string `json:"last_name,omitempty"`
	Username        string `json:"username,omitempty"`
	Profile_picture string `json:"profile_picture,omitempty"`
}

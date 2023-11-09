package resources

type UserResource struct {
	Id              int    `json:"id,omitempty"`
	First_name      string `json:"first_name,omitempty"`
	Last_name       string `json:"last_name,omitempty"`
	Email           string `json:"email,omitempty"`
	Username        string `json:"username,omitempty"`
	Profile_picture string `json:"profile_picture,omitempty"`
	Phone_number    int    `json:"phone_number,omitempty"`
	Address         string `json:"address,omitempty"`
	Status          string `json:"status,omitempty"`
}

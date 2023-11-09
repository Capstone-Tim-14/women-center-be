package resources

type AuthResource struct {
	Id       uint   `json:"id,omitempty"`
	Fullname string `json:"fullname,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
}

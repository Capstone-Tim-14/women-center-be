package resources

type AdminResource struct {
	Id         int           `json:"id,omitempty"`
	First_name string        `json:"first_name,omitempty"`
	Last_name  string        `json:"last_name,omitempty"`
	Email      string        `json:"email,omitempty"`
	Username   string        `json:"username,omitempty"`
	Role       *RoleResource `json:"role,omitempty"`
}

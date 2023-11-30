package resources

type CounselorResource struct {
	Id                uint                  `json:"id,omitempty"`
	First_name        string                `json:"first_name,omitempty"`
	Last_name         string                `json:"last_name,omitempty"`
	Email             string                `json:"email,omitempty"`
	Username          string                `json:"username,omitempty"`
	Profile_picture   string                `json:"profile_picture,omitempty"`
	Phone_number      string                `json:"phone_number,omitempty"`
	Description       string                `json:"description,omitempty"`
	Address           string                `json:"address,omitempty"`
	Certification_url string                `json:"certification,omitempty"`
	Status            string                `json:"status,omitempty"`
	Specialist        []SpecialistCounselor `json:"specialist,omitempty"`
}

type SpecialistCounselor struct {
	Id   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

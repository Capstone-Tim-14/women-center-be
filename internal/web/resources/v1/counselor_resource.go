package resources

type CounselorResource struct {
	Id                uint                  `json:"id,omitempty"`
	First_name        string                `json:"first_name,omitempty"`
	Last_name         string                `json:"last_name,omitempty"`
	Email             string                `json:"email,omitempty"`
	Username          string                `json:"username,omitempty"`
	Profile_picture   string                `json:"profile_picture,omitempty"`
	Education         string                `json:"education,omitempty"`
	Phone_number      string                `json:"phone_number,omitempty"`
	Description       string                `json:"description,omitempty"`
	Address           string                `json:"address,omitempty"`
	Certification_url string                `json:"certification,omitempty"`
	Status            string                `json:"status,omitempty"`
	Specialist        []SpecialistCounselor `json:"specialist,omitempty"`
}

type ProfileCounselor struct {
	Id              uint   `json:"id,omitempty"`
	Profile_picture string `json:"profile_picture,omitempty"`
	Username        string `json:"username,omitempty"`
	Full_name       string `json:"full_name,omitempty"`
	Email           string `json:"email,omitempty"`
}

type SpecialistCounselor struct {
	Id   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CounselingSchedule struct {
	Day_schedule  string                   `json:"day_schedule,omitempty"`
	Time_schedule []CounselingScheduleTime `json:"time_schedule,omitempty"`
}

type CounselingScheduleTime struct {
	Time_start  string `json:"time_start,omitempty"`
	Time_finish string `json:"time_finish,omitempty"`
}

type DetailCounselor struct {
	Id              string                `json:"id,omitempty"`
	Full_name       string                `json:"full_name,omitempty"`
	Email           string                `json:"email,omitempty"`
	Usernam         string                `json:"username,omitempty"`
	Profile_picture string                `json:"profile_picture,omitempty"`
	Phone_number    string                `json:"phone_number,omitempty"`
	Description     string                `json:"description,omitempty"`
	Address         string                `json:"address,omitempty"`
	Spesialist      []SpecialistCounselor `json:"specialist,omitempty"`
	Schedule        []CounselingSchedule  `json:"schedule,omitempty"`
}

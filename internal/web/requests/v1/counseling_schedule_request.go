package requests

type CounselingScheduleRequest struct {
	Day_schedule string `json:"day_schedule" validate:"required" form:"day_schedule"`
	Time_start   string `json:"time_start" validate:"required,ltefield=Time_finish" form:"time_start"`
	Time_finish  string `json:"time_finish" validate:"required" form:"time_finish"`
}

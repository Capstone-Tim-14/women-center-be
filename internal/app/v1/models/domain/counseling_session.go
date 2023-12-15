package domain

type CounselingSession struct {
	OrderId       string
	User_id       uint
	First_name    string
	Last_name     string
	Package_title string
	Date_schedule string
	Status        string
	Email         string
	Day_schedule  string
	Time_starts   string
	Time_finishs  string
	Time_start    string
	Time_finish   string
}

type CounselingSessionDetail struct {
	OrderId       string
	First_name    string
	Last_name     string
	Package_title string
	Status        string
	Email         string
}

type CounselingScheduleSession struct {
	Date_schedule string
	Time_start    string
	Time_finish   string
}

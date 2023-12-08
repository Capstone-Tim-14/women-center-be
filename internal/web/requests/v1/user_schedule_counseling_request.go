package requests

import "time"

type UserScheduleCounselingQueryRequest struct {
	Counselor_schedule_id uint
	Day                   time.Time
	Time_start            string
}

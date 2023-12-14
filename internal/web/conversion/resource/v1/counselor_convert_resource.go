package conversion

import (
	"strings"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/utils/helpers"
)

func CounselorDomainToCounselorResponse(counselor *domain.Counselors) resources.CounselorResource {
	return resources.CounselorResource{
		Id:              counselor.Id,
		First_name:      counselor.First_name,
		Last_name:       counselor.Last_name,
		Email:           counselor.Credential.Email,
		Description:     counselor.Description,
		Status:          counselor.Status,
		Profile_picture: counselor.Profile_picture,
	}
}

func ConvertCounselorDomainToCounselorResponse(counselor []domain.Counselors) []resources.CounselorResource {
	counselorResponse := []resources.CounselorResource{}

	for _, counselor := range counselor {
		counselorResponse = append(counselorResponse, resources.CounselorResource{
			Id:              counselor.Id,
			First_name:      counselor.First_name,
			Last_name:       counselor.Last_name,
			Email:           counselor.Credential.Email,
			Education:       counselor.Education,
			Description:     counselor.Description,
			Status:          counselor.Status,
			Profile_picture: counselor.Profile_picture,
		})
	}

	return counselorResponse
}

func ConvertCounselorDomainToCounselorDetailResponse(counselor *domain.Counselors, schedules []domain.Counseling_Schedule) resources.DetailCounselor {
	counselorResponse := resources.DetailCounselor{
		Full_name:       counselor.First_name + " " + counselor.Last_name,
		Email:           counselor.Credential.Email,
		Description:     counselor.Description,
		Profile_picture: counselor.Profile_picture,
	}

	for _, specialist := range counselor.Specialists {
		counselorResponse.Spesialist = append(counselorResponse.Spesialist, resources.SpecialistCounselor{
			Id:   specialist.Id,
			Name: specialist.Name,
		})
	}

	for _, schedule := range schedules {
		var CounselingSchedule resources.CounselingSchedule
		var TimeSchedule []resources.CounselingScheduleTime

		CounselingSchedule = resources.CounselingSchedule{
			Day_schedule: schedule.Day_schedule,
		}

		TimeStartSplit := strings.Split(*schedule.Time_starts, ",")
		TimeEndSplit := strings.Split(*schedule.Time_finishs, ",")

		TimeSchedulePerTime := resources.CounselingScheduleTime{}

		for i, timeStart := range TimeStartSplit {
			timeStrs := strings.Trim(timeStart, "[]")

			if i < len(TimeEndSplit) {

				timeEndStrs := strings.Trim(TimeEndSplit[i], "[]")

				TimeSchedulePerTime.Time_start = helpers.ParseTimeToClock(helpers.ParseStringToTime(timeStrs))
				TimeSchedulePerTime.Time_finish = helpers.ParseTimeToClock(helpers.ParseStringToTime(timeEndStrs))
			}
			TimeSchedule = append(TimeSchedule, TimeSchedulePerTime)
		}

		CounselingSchedule.Time_schedule = TimeSchedule

		counselorResponse.Schedule = append(counselorResponse.Schedule, CounselingSchedule)

	}

	return counselorResponse
}

func CounselorDomainToProfileCounselor(counselor *domain.Counselors) resources.ProfileCounselor {
	return resources.ProfileCounselor{
		Id:              counselor.Id,
		Profile_picture: counselor.Profile_picture,
		Username:        counselor.Credential.Username,
		Full_name:       counselor.First_name + " " + counselor.Last_name,
		Email:           counselor.Credential.Email,
	}
}

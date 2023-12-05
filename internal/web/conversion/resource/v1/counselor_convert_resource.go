package conversion

import (
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
			Description:     counselor.Description,
			Status:          counselor.Status,
			Profile_picture: counselor.Profile_picture,
		})
	}

	return counselorResponse
}

func ConvertCounselorDomainToCounselorDetailResponse(counselor *domain.Counselors) resources.DetailCounselor {
	counselorResponse := resources.DetailCounselor{
		First_name:      counselor.First_name,
		Last_name:       counselor.Last_name,
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

	for _, schedule := range counselor.Schedules {
		counselorResponse.Schedule = append(counselorResponse.Schedule, resources.CounselingSchedule{
			Day_schedule: schedule.Day_schedule,
			Time_start:   helpers.ParseTimeToClock(&schedule.Time_start),
			Time_finish:  helpers.ParseTimeToClock(&schedule.Time_finish),
		})
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

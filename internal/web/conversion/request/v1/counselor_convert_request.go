package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/helpers"
)

func CounselorCreateRequestToCounselorDomain(request requests.CounselorRequest) *domain.Counselors {
	return &domain.Counselors{
		First_name: request.First_name,
		Last_name:  request.Last_name,
		Credential: &domain.Credentials{
			Username: request.Username,
			Password: request.Password,
			Role_id:  request.Role_id,
			Email:    request.Email,
		},
		Profile_picture: request.Profile_picture,
		Description:     request.Description,
		Status:          "ACTIVE",
	}
}

func CounselorUpdateRequestToCounselorDomainForMobile(request requests.UpdateCounselorProfileRequestForMobile, counselor *domain.Counselors) *domain.Counselors {
	counselor.First_name = request.First_name
	counselor.Last_name = request.Last_name
	counselor.Credential.Role_id = request.Role_id
	counselor.Credential.Email = request.Email
	counselor.Profile_picture = request.Profile_picture
	counselor.Birthday = helpers.ParseStringToTime(request.Birthday)

	return counselor
}

func CounselorUpdateRequestToCounselorDomain(request requests.UpdateCounselorProfileRequest, counselor *domain.Counselors) *domain.Counselors {
	counselor.First_name = request.First_name
	counselor.Last_name = request.Last_name
	counselor.Credential.Role_id = request.Role_id
	counselor.Credential.Email = request.Email
	counselor.Profile_picture = request.Profile_picture
	counselor.Description = request.Description
	counselor.Education = request.Education
	counselor.Birthday = helpers.ParseStringToTime(request.Birthday)

	return counselor
}

func CounselorScheduleToCounselorDomain(requests []requests.CounselingScheduleRequest) []domain.Counseling_Schedule {

	var CounselorScheduling []domain.Counseling_Schedule

	for _, val := range requests {

		Scheduling := domain.Counseling_Schedule{
			Day_schedule: val.Day_schedule,
			Time_start:   helpers.ParseClockToTime(val.Time_start),
			Time_starts:  nil,
			Time_finishs: nil,
			Time_finish:  helpers.ParseClockToTime(val.Time_finish),
		}

		CounselorScheduling = append(CounselorScheduling, Scheduling)
	}

	return CounselorScheduling

}

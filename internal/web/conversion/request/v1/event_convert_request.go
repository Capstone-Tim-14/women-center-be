package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/helpers"
)

func EventCreateRequestToEventDomain(request *requests.EventRequest) *domain.Event {
	return &domain.Event{
		Title:       request.Title,
		Poster:      *request.Poster,
		Locations:   request.Locations,
		Date:        helpers.ParseStringToTime(request.Date),
		Price:       helpers.StringToDecimal(request.Price),
		Time_start:  helpers.ParseClockToTime(request.Time_start),
		Time_finish: helpers.ParseClockToTime(request.Time_finish),
		EventUrl:    request.EventUrl,
		EventType:   request.EventType,
		Status:      request.Status,
		Description: request.Description,
	}
}

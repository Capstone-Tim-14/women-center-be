package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/utils/helpers"
)

func EventDetailDomainToEventResource(event *domain.Event) *resources.EventResource {
	return &resources.EventResource{
		Id:          event.Id,
		Title:       event.Title,
		Description: event.Description,
		Poster:      event.Poster,
		Locations:   event.Locations,
		Date:        helpers.ParseOnlyDate(event.Date),
		Price:       int(event.Price.IntPart()),
		Time_start:  helpers.ParseTimeToClock(&event.Time_start),
		Time_finish: helpers.ParseTimeToClock(&event.Time_finish),
		EventUrl:    event.EventUrl,
		EventType:   event.EventType,
		Status:      event.Status,
		CreatedAt:   helpers.ParseOnlyDate(&event.CreatedAt),
		UpdatedAt:   helpers.ParseOnlyDate(&event.UpdatedAt),
	}
}

func AllEventConvertResource(events []domain.Event) []resources.AllEventResource {
	eventResource := []resources.AllEventResource{}
	for _, event := range events {
		singleEventResource := resources.AllEventResource{}
		singleEventResource.Id = event.Id
		singleEventResource.Title = event.Title
		singleEventResource.Locations = event.Locations
		singleEventResource.Date = helpers.ParseOnlyDate(event.Date)
		singleEventResource.Time_start = helpers.ParseTimeToClock(&event.Time_start)
		singleEventResource.Time_finish = helpers.ParseTimeToClock(&event.Time_finish)
		eventResource = append(eventResource, singleEventResource)
	}

	return eventResource
}

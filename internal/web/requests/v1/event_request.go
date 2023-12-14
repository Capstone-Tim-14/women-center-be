package requests

type EventRequest struct {
	Title       string  `json:"title" validate:"required" form:"title" form:"title"`
	Poster      *string `json:"poster" validate:"required" form:"poster" form:"poster"`
	Locations   string  `json:"locations" validate:"required" form:"locations" form:"locations"`
	Date        string  `json:"date" validate:"required" form:"date" form:"date"`
	Price       string  `json:"price" validate:"required" form:"price" form:"price"`
	EventUrl    string  `json:"event_url" validate:"required" form:"event_url" form:"event_url"`
	Time_start  string  `json:"time_start" validate:"required,ltefield=Time_finish" form:"time_start"`
	Time_finish string  `json:"time_finish" validate:"required" form:"time_finish"`
	EventType   string  `json:"event_type" validate:"required" form:"event_type" form:"event_type"`
	Status      string  `json:"status" validate:"required" form:"status" form:"status"`
	Description string  `json:"description" validate:"required" form:"description" form:"description"`
}

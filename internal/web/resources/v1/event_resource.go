package resources

type EventResource struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Poster      string `json:"poster"`
	Locations   string `json:"locations"`
	Date        string `json:"date"`
	Price       int    `json:"price"`
	Time_start  string `json:"time_start"`
	Time_finish string `json:"time_finish"`
	EventUrl    string `json:"event_url"`
	EventType   string `json:"event_type"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

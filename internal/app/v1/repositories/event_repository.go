package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(event *domain.Event) (*domain.Event, error)
}

type EventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &EventRepositoryImpl{
		db: db,
	}
}

func (repository *EventRepositoryImpl) CreateEvent(event *domain.Event) (*domain.Event, error) {
	result := repository.db.Create(&event)
	if result.Error != nil {
		return nil, result.Error
	}

	return event, nil
}

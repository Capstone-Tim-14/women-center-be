package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type EventRepository interface {
	CreateEvent(event *domain.Event) (*domain.Event, error)
	FindDetailEvent(id int) (*domain.Event, error)
	FindAllEvent() ([]domain.Event, error)
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

func (repository *EventRepositoryImpl) FindDetailEvent(id int) (*domain.Event, error) {
	var event domain.Event
	result := repository.db.Where("id = ?", id).First(&event)
	if result.Error != nil {
		return nil, fmt.Errorf("Event not found")
	}

	return &event, nil
}

func (repository *EventRepositoryImpl) FindAllEvent() ([]domain.Event, error) {
	var events []domain.Event
	result := repository.db.Find(&events)
	if result.Error != nil {
		return nil, fmt.Errorf("Event not found")
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Event empty")
	}

	return events, nil
}

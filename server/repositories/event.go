package repositories

import (
	"context"
	"time"

	"github.com/Malayt04/BookTicket/backend/models"
)

type EventRepository struct {
	db any
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := [] *models.Event{}

	events = append(events, &models.Event{
		ID: 1,
		Name: "Event 1",
		Location: "Location 1",
		Date: "2023-01-01",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, id int) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) error {
	return nil
}

func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{db}
}
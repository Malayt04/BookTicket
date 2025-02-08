package repositories

import (
	"context"

	"github.com/Malayt04/BookTicket/backend/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := [] *models.Event{}

	res := r.db.Model(&models.Event{}).Find(&events)

	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, id uint) (*models.Event, error) {
	event := &models.Event{}

	res := r.db.Model(&models.Event{}).Where("id = ?", id).First(&event)

	if res.Error != nil{
		return nil, res.Error
	}

	return event, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) error {
	
	res := r.db.Create(event)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func(r *EventRepository) UpdateEvent(ctx context.Context, event *models.Event) error {

	res := r.db.Model(&models.Event{}).Where("id = ?", event.ID).Updates(event)

	if res.Error != nil {
		return res.Error
	}

	return nil

}

func (r *EventRepository) DeleteEvent(ctx context.Context, id uint)error{

	res := r.db.Model(&models.Event{}).Where("id = ?", id).Delete(&models.Event{})

	if res.Error != nil {
		return res.Error
	}

	return nil

}

func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
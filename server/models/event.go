package models

import (
	"time"
	"context"
)

type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Location string `json:"location"`
	Date string `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EventRepository interface {
	GetMany(ctx context.Context) ([]*Event, error)
	GetOne(ctx context.Context, id int) (*Event, error)
	CreateOne(ctx context.Context, event *Event) error
}
package models

import (
	"context"
	"time"
)

type Ticket struct {
	ID        uint      `json:"id"`
	EventID   uint      `json:"event_id"`
	Event Event `json:"event"`
	Entered bool `json:"entered"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TicketRepository interface {

	GetMany(ctx context.Context) ([]*Ticket, error)
	GetOne(ctx context.Context, ticket_id int)(*Ticket, error)
	CreateTicket(ctx context.Context, ticket *Ticket) error
}
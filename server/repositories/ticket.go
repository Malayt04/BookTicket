package repositories

import (
	"context"

	"github.com/Malayt04/BookTicket/backend/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func(r  *TicketRepository) GetMany(ctx context.Context) ([]*models.Ticket, error){
	
	tickets := [] *models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Find(&tickets)

	if res.Error != nil {
		return nil, res.Error
	}

	return tickets, nil
	
}

func(r  *TicketRepository) GetOne(ctx context.Context, id uint)(*models.Ticket, error){

	ticket := &models.Ticket{}

	res := r.db.Model((&models.Ticket{})).Where("id = ?", id).Find(&ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func(r  *TicketRepository) CreateTicket(ctx context.Context, ticket *models.Ticket) error{

	res := r.db.Model(&models.Ticket{}).Create(ticket)

	if res.Error != nil{
		return res.Error
	}

	return nil

}

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{
		db: db,
	}
}
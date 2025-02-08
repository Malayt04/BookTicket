package handlers

import (
	"github.com/Malayt04/BookTicket/backend/models"
	"github.com/gofiber/fiber/v2"
)

type TicketHandler struct{
	repository models.TicketRepository
}

func (h *TicketHandler) GetMany(c *fiber.Ctx){

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tickets, err := h.repository.GetMany(context)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch tickets",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"tickets": tickets,
	})

}

func (h *TicketHandler) CreateOne(c *fiber.Ctx){

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ticket := make(models.Ticket)

	if err := c.BodyParser(&ticket); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to parse request body",
		})
	}

	if err := h.repository.Create(context, ticket); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create ticket",
		})
	}

}

func (h *TicketHandler) GetOne(c *fiber.Ctx){

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, _ := strcnv(c.Params("id"))

	ticket, err := h.repository.GetOne(context, uint(id))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Ticket not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ticket": ticket,
	})

}

func NewTicketHandler(router *fiber.Router, repository *,models.TicketRepository) {
	handler := &TicketHandler{
		repository: *repository,
	}

	(*router).Get("/tickets", handler.GetMany)
	(*router).Get("/tickets/:id", handler.GetOne)
	(*router).Post("/tickets", handler.CreateOne)
}
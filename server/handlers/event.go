package handlers

import (
	"context"
	"time"

	"github.com/Malayt04/BookTicket/backend/models"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	repository models.EventRepository
}

func (h *EventHandler) GetMany(c *fiber.Ctx) error {

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	events, err := h.repository.GetMany(context)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch events",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"events": events,
	})

}

func (h *EventHandler) GetOne(c *fiber.Ctx) error {
	return nil
}

func (h *EventHandler) CreateOne(c *fiber.Ctx) error {
	return nil
}

func NewEventHandler(router fiber.Router, repository models.EventRepository){
	handler := &EventHandler{repository: repository}

	router.Get("/events", handler.GetMany)
	router.Get("/events/:id", handler.GetOne)
	router.Post("/events", handler.CreateOne)
} 
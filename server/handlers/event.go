package handlers

import (
	"context"
	"strconv"
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
	
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, _ := strconv.Atoi(c.Params("id"))



	event, err := h.repository.GetOne(context, uint(id))

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Event not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"event": event,
	})

}

func (h *EventHandler) CreateOne(c *fiber.Ctx) error {
	
	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	event := new(models.Event)

	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
		})
	}

	if err := h.repository.CreateOne(context, event); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create event",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"event": event,
	})

}

func (h *EventHandler) UpdateEvent(c *fiber.Ctx) error {

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	event := new(models.Event)

	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse request body",
		})
	}

	if err := h.repository.UpdateEvent(context,event); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update event",
		})
	} 

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"event": event,
	})

}

func (h *EventHandler) DeleteEvent(c *fiber.Ctx) error {

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id, _ := strconv.Atoi(c.Params("id"))

	if err := h.repository.DeleteEvent(context, uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to delete event",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message":"Event deleted successfully",
	})

}

func NewEventHandler(router fiber.Router, repository models.EventRepository){
	handler := &EventHandler{repository: repository}

	router.Get("/", handler.GetMany)
	router.Get("/:id", handler.GetOne)
	router.Post("/", handler.CreateOne)
	router.Put("/:id", handler.UpdateEvent)
	router.Delete("/", handler.DeleteEvent)
} 
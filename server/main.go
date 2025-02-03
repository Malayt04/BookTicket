package main

import (
	"github.com/Malayt04/BookTicket/backend/handlers"
	"github.com/Malayt04/BookTicket/backend/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		AppName: "BookTicket",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(nil)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":8080")

}
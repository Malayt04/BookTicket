package main

import (
	"github.com/Malayt04/BookTicket/backend/config"
	"github.com/Malayt04/BookTicket/backend/db"
	"github.com/Malayt04/BookTicket/backend/handlers"
	"github.com/Malayt04/BookTicket/backend/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {

	newConfig := config.NewEnvConfig()
	db := db.Init(newConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName: "BookTicket",
		ServerHeader: "Fiber",
	})

	eventRepository := repositories.NewEventRepository(db)

	server := app.Group("/api")

	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(config.NewEnvConfig().SERVER_PORT)

}
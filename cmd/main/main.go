package main

import (
	"log"

	"github.com/amanuel15/fiber_server/pkg/configs"
	"github.com/amanuel15/fiber_server/pkg/database"
	routes "github.com/amanuel15/fiber_server/pkg/routes"
	"github.com/amanuel15/fiber_server/pkg/websocket"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.SetupEnvironmentVariables()
	database.ConnectDB()
	app := fiber.New()
	websocket.Websocket(app)
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":" + configs.PORT))
}

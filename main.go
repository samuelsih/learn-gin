package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/samuelsih/learn-gofiber/routes"
)

const port = ":8080"

var (
	app *fiber.App
)

func main() {
	app = fiber.New()

	SetMiddeware(app)

	routes.Router(app)

	log.Fatal(app.Listen(port))
}
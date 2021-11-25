package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samuelsih/learn-gofiber/handler"
)

//this is where routes are store
func Router(app *fiber.App) {
	app.Get("/", handler.HomePage)
	app.Get("/:param1/:param2/:param3", handler.ThreeParameter)
	app.Get("/*", handler.RandomParameter)
	
}

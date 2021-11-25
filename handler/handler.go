package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

)

type Response struct {
	Res string `json:"response"`
}

//HomePage returns string Hello from homepage
func HomePage(c *fiber.Ctx) error {
	return c.SendString("Hello from homepage")
}

//RandomParameter returns string you type in parameter
//example = http://localhost:8080/test123
//it will return "You typed test123"
func RandomParameter(c *fiber.Ctx) error {
	param := c.Params("*")
	
	message := fmt.Sprintf("You typed %s", param)

	return c.SendString(message)
}

//ThreeParameter return string you type in http://localhost:8080/param1/param2/param3
func ThreeParameter(c *fiber.Ctx) error {
	param1 := c.Params("param1")
	param2 := c.Params("param2")
	param3 := c.Params("param3")

	message := fmt.Sprintf("You guys wrote 3 parameters, %s, %s and %s. Horray", param1, param2, param3)
	
	return c.SendString(message)
}
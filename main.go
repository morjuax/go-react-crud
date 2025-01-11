package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/morjuax/go-react-crud/utils"
)

func main() {
	app := fiber.New()

	app.Listen(":3000")
	utils.Print("Server on port 3000")
}
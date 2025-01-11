package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/morjuax/go-react-crud/utils"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	} 

	app := fiber.New()

	app.Static("/", "./client/dist")

	app.Use(cors.New())

	app.Get("/users", func (c *fiber.Ctx) error  {
		return c.JSON(&fiber.Map{
			"data": "users from backend",
		})
	})

	app.Listen(":" + port)
	utils.Print("Server on port 3000")
}
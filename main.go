package main

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/morjuax/go-react-crud/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	} 

	app := fiber.New()
	
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/gomongodb"))

	if err != nil {
		panic(err)
	}

	coll := client.Database("gomongodb").Collection("users")
	coll.InsertOne(context.TODO(), bson.D{{
		Key: "name", 
		Value: "morjuax",
	}})

	app.Static("/", "./client/dist")

	app.Use(cors.New())

	app.Get("/users", func (c *fiber.Ctx) error  {
		return c.JSON(&fiber.Map{
			"data": "users from backend",
		})
	})

	app.Post("/users", func (c *fiber.Ctx) error  {
		var user
		return c.JSON(&fiber.Map{
			"data": "creating user",
		})
	})

	app.Listen(":" + port)
	utils.Print("Server on port 3000")
}
package main

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/morjuax/go-react-crud/models"
	"github.com/morjuax/go-react-crud/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        utils.Print("Error cargando archivo .env: ", err)
    }


	port := os.Getenv("PORT")
	mongodb := os.Getenv("MONGODB_URI")
	
	if port == "" {
		port = "3000"
	} 
	
	if mongodb == "" {
		mongodb = "mongodb://localhost:27017/gomongodb"
	}

	app := fiber.New()
	
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongodb))

	if err != nil {
		panic(err)
	}

	app.Static("/", "./client/dist")

	app.Use(cors.New())

	// app.Get("/users", func (c *fiber.Ctx) error  {
	// 	return c.JSON(&fiber.Map{
	// 		"data": "users from backend",
	// 	})
	// })

	app.Post("/users", func (c *fiber.Ctx) error  {
		var user models.User

		c.BodyParser(&user)

		
		coll := client.Database("gomongodb").Collection("users")
		result, err := coll.InsertOne(context.TODO(), bson.D{{
			Key: "name", 
			Value: user.Name,
		}})

		if err != nil {
			panic(err)
		}


		return c.JSON(&fiber.Map{
			"data": result,
		})
	})

	app.Get("/users", func (c *fiber.Ctx) error  {
		var users []models.User

		coll := client.Database("gomongodb").Collection("users")
		results, err := coll.Find(context.TODO(), bson.M{})

		if err != nil {
			panic(err)
		}
		for results.Next(context.TODO()) {
			var user models.User
			results.Decode(&user)
			users = append(users, user)
		}
		return c.JSON(&fiber.Map{
			"users": users,
		})
	})

	app.Listen(":" + port)
	utils.Print("Server on port 3000")
}
package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}


	app:= fiber.New()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/gomongodb")) //  conecto base de datos

	if err != nil {
		panic(err) //acabo la ejecucion del programa y me muestra el error
	}

	coll := client.Database("gomongodb").Collection("users")
	coll.InsertOne(context.TODO(), bson.D{{
		Key: "name",
		Value:"jose",
	}})

	app.Use(cors.New())

	app.Static("/", "./veterinary-client/dist")

	app.Get("/users", func(c *fiber.Ctx ) error {
		return c.JSON(&fiber.Map{
			"data": "usuarios desde el backend",
		})
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "creando usuario",
		})
	})

	app.Listen(":3000")

	fmt.Println("server on port 3000")


}
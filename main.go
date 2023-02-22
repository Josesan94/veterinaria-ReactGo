package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"

	"github.com/Josesan94/veterinaria-ReactGo/routes"
)


// func setupRoutes(app *fiber.App) {

// 	app.Post("/api/users", routes.CreateUser)
// 	app.Get("/api/users", routes.GetUsers)
// }

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}


	app:= fiber.New()
	// setupRoutes(app)

	app.Use(cors.New())

	

	app.Static("/", "./veterinary-client/dist")



	routes.CreateUser(app)
	routes.GetUsers(app)

	// app.Get("/users", func(c *fiber.Ctx ) error {
	// 	var users []models.User

	// 	coll := client.Database("gomongodb").Collection("users")
	// 	results, error := coll.Find(context.TODO(), bson.M{}) //recorro la data de results, lo decodifico y lo añado al usuario (linea 67)

	// 	if error != nil {
	// 		panic(error)
	// 	}

	// 	for results.Next(context.TODO()) {
	// 		var user models.User
	// 		results.Decode(&user)
	// 		users = append(users, user)
	// 	}

	// 	return c.JSON(&fiber.Map{
	// 		"users": users,
	// 	})

	// })

	// func findUser(id int, user *models.User) error {

	// }

	// app.Get("/user", func(c *fiber.Ctx) error {
	// 	id, err := c.ParamsInt("id")

	// 	var user models.User

	// 	if err != nil  {
	// 		return c.Status(400).JSON("asegurese que el id es un numero")
	// 	}


	// })

	app.Listen(":3000")

	fmt.Println("server on port 3000")


}
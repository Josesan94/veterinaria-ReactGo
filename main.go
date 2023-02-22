package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"github.com/Josesan94/veterinaria-ReactGo/configs"
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

	 //run database
	 configs.ConnectDB()

	app.Static("/", "./veterinary-client/dist")



	routes.CreateUser(app)
	routes.GetUsers(app)


	app.Listen(":3000")

	fmt.Println("server on port 3000")


}
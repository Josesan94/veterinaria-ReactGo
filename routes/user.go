package routes

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gofiber/fiber/v2"
	"github.com/Josesan94/veterinaria-ReactGo/controllers"
)

//serializer del usuario
type User struct {
	ID     primitive.ObjectID  `json:"id" bson:"_id"`
	Name 	string	`json:"name"`
}



func UserRoute(app *fiber.App) {
    app.Post("/user", controllers.CreateUser)
	app.Get("/users", controllers.GetAllUsers)
	app.Get("/user/:userId", controllers.FindUser)
	app.Put("/user/:userId", controllers.UpdateUser)
	app.Delete("/user/:userId", controllers.DeleteUser)
}



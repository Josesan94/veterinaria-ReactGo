package routes

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/gofiber/fiber/v2"
	"github.com/Josesan94/veterinaria-ReactGo/models"
	"time"

)

//serializer del usuario
type User struct {
	ID     primitive.ObjectID  `json:"id" bson:"_id"`
	Name 	string	`json:"name"`
}



func CreateUser(router fiber.Router) {
	router.Post("/users", func(c *fiber.Ctx) error {
		var user models.User

		c.BodyParser(&user) // me extrae los valores que me envia el frontend
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/gomongodb")) //  conecto base de datos

		coll := client.Database("gomongodb").Collection("users")
		coll.InsertOne(context.TODO(), bson.D{{
		Key: "name",
		Value: user.Name,
	}})

	if err != nil {
		panic(err)
	}

		return c.JSON(&fiber.Map{
			"user": user.Name,
		})
	})

}

func GetUsers(router fiber.Router) {
	var users []models.User
	var user models.User
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017/gomongodb")) //  conecto base de datos
	coll := client.Database("gomongodb").Collection("users")
	

	router.Get("/users", func(c *fiber.Ctx ) error {

		results, error := coll.Find(context.TODO(), bson.M{}) //recorro la data de results, lo decodifico y lo añado al usuario (linea 67)

		if err != nil {
			panic(error)
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

	router.Get("/users/:id", func(c *fiber.Ctx) error { 

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Params("id")
		coll2 := client.Database("gomongodb").Collection("users")
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(id)

		err := coll2.FindOne(ctx, bson.M{"id": objId}).Decode(&user)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}


		return c.JSON(&fiber.Map {
			"user": user,
		})
	})

}




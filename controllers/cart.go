package controllers

import (
	"github.com/Josesan94/veterinaria-ReactGo/configs"
	"time"
	"context"
	"errors"
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
)


type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}


func (app *Application) AddToCart(c *fiber.Ctx)  {
	productQueryID := c.Query("id")
	if productQueryID == "" {

		c.Status(http.StatusBadRequest).JSON("product id is empty")
		return
	}

	userQueryID := c.Query("userId")

	if	userQueryID == "" {
		_ = c.Status(http.StatusBadRequest).JSON("user id is empty")
		return
	}

	productID, err := primitive.ObjectIDFromHex(productQueryID)

	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError).JSON("product id is empty")
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	configs.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID )

	if err != nil {
		return c.IndentedJSON(http.StatusInternalServer, err)
	}
	c.IndentedJSON(200, "Succesfully added to the cart")

}

func (app *Application) RemoveItem(c *fiber.Ctx) {
	productQueryID := c.Query("id")
	if productQueryID == "" {
		log.Println("product id is empty")

		return c.abortWithError(http.StatusBadRequest, errors.New("product id is empty"))
	}

	userQueryID := c.Query("userId")

	if	userQueryID == "" {
		_ = c.abortWithError(http.StatusBadRequest, errors.New("user id is empty"))
		return
	}

	productID, err := primitive.ObjectIdFromHex(productQueryID)

	if err != nil {
		log.Println(err)
		c.AbourtWithStatus(http.StatusInternalServerError)
		return
	}

	var ctx, cancel = context.withTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = configs.RemoveCartItem(ctx, app.prodCollection, app.userCollection, productID, userQueryID)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServer, err)
	}

	c.IndentedJSON(200, "Succesfully removed item from the cart")
}

func GetCart() {
	
}
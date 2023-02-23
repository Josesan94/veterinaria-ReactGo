package controllers

import (
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


func (app *Application) AddToCart(c *fiber.Context)  {
	productQueryID := c.Query("id")
	if productQueryid == "" {
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

	configs.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID )

	if err != nil {
		return c.IndentedJSON(http.StatusInternalServer, err)
	}
	c.IndentedJSON(200, "Succesfully added to the cart")

}

func (app *Application) RemoveItem(c *fiber.Context) {
	productQueryID := c.Query("id")
	if productQueryid == "" {
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

func SeeCart() {

}
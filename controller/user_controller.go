package controllers

import (
    "context"
    "github.com/Josesan94/veterinaria-ReactGo/configs"
    "github.com/Josesan94/veterinaria-ReactGo/models"
    "github.com/Josesan94/veterinaria-ReactGo/responses"
    "net/http"
    "time"

    "github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()
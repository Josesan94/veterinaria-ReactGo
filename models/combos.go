package models


import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CatCombo struct {
	ID 				primitive.ObjectID `bson:"id"`
	Combo		*string   			`json:combo`

}

type DogCombo struct {
	ID 				primitive.ObjectID `bson:"id"`
	Combo		*string   			`json:combo`

}
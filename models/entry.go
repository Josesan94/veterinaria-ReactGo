package models


import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CatCombo struct {
	Id 				primitive.ObjectID `bson:"id"`
	Quantity		*string   			`json:quantity`
	Complementary1	*string				`json:complementary1`
	Complementary2	*string				`json:complementary2`

}

type DogCombo struct {
	Id 				primitive.ObjectID `bson:"id"`
	Quantity		*string   			`json:quantity`
	Complementary1	*string				`json:complementary1`
	Complementary2	*string				`json:complementary2`
}
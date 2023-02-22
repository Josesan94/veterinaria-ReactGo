package models

import (

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     			primitive.ObjectID `json:"id,omitempty"`
	Name  			string `json:"name,omitempty" validate:"required"`
	Email			string `json:"email" validate:"required"`
	Password		string `json:"password validate:"required"`
	Token			string `json:"token"`
	Refresh_Token	string `json:"refresh_token"`
	User_ID		string `json:"user_id"`
	Address			string `json:"address"`
	UserCart		[]ProductUser `json:"usercart"`
}

type ProductUser struct {

}

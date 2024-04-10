package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
// model in DB
type Entry struct { 
	ID			primitive.ObjectID	`bson:"id"` // ID created by GO so we dont have to pass it help of bson pkg Golang understands this type
	Dish		*string				`json: "dish"` // *string passes refernece
	Fat			*string				`json: "fat"`
	Ingredients	*string				`json: "ingredients"`
	Calories	*string				`json: "calories"`
}
// marshalling and unmarshalling converting from json to structs and visa versa
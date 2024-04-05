package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
// model in DB
type Entry struct { 
	ID			primitive.ObjectID	`bson:"id` // ID created by GO so we dont have to pass it
	Dish		*string				`json: "dish"` // *string passes refernece
	Fat			*float64			`json: "fat"`
	Ingredients	*string				`json: "indgredients"`
	Calories	*string				`json: "calories"`
}
// marshalling and unmarshalling converting from json to structs and visa versa
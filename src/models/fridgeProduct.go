package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type fridgeProduct struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty"`
	Category string             `bson:"category,omitempty"`
	InsertTS primitive.DateTime `bson:"insertts,omitempty"`
}

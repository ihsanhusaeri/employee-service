package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

//Employee representation data
type Employee struct {
	ID       *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string              `json:"name" bson:"name"`
	Address  string              `json:"address" bson:"address"`
	Position string              `json:"position" bson:"position"`
}

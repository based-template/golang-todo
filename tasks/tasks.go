package tasks

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task : the base data structure for the todo app
type Task struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Item      string             `json:"item,omitempty"`
	Completed bool               `json:"completed,omitempty"`
}

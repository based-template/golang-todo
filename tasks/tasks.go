package tasks

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Task : the base data structure for the todo app
type Task struct {
	ID        primitive.ObjectID
	Completed bool
	Item      string
	Date      time.Date
}

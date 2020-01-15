package tasks

import "time"

// Task : the base data structure for the todo app
type Task struct {
	Completed bool
	Item      string
	Date      time.Date
}

package router

import (
	"github.com/gorilla/mux"
	"golang-todo/services"
)

// Router : setup handlers for API routes
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/tasks", services.GetTaskList).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", services.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", services.GetTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", services.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}/complete", services.MarkCompleted).Methods("PUT", "OPTIONS")
	return router
}

package router

import (
	"../services"
	"github.com/gorilla/mux"
)

// Router : setup handlers for API routes
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/tasks", services.GetTaskList).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", services.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", services.GetTask).Methods("GET", "OPTIONS")
	return router
}

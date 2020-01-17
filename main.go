package main

import (
	"fmt"
	"log"
	"net/http"

	"golang-todo/router"
)

func main() {
	r := router.Router()
	fmt.Println("Launching server")

	log.Fatal(http.ListenAndServe(":8080", r))
}

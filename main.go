package main

import (
	"audio-crud-nooble/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	fmt.Println("Server started on the port 8080!")
	log.Fatal(http.ListenAndServe(":8080", r))
}

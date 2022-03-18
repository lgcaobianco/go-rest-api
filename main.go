package main

import (
	"go-rest-api/database"
	"go-rest-api/routes"
	"log"
)

func main() {
	database.Connect()
	log.Println("project running")
	routes.HandleRequest()
}
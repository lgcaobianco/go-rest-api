package main

import (
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {
	models.Products = []models.Product{
		{Id: 1, Name: "iphone x", EditorialName: "iphone X 128gb"},
		{Id: 2, Name: "samsung ", EditorialName: "iphone X 128gb"},
	}
	routes.HandleRequest()
}
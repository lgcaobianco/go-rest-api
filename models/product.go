package models

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	EditorialName string `json:"editorial_name"`
}

var Products []Product
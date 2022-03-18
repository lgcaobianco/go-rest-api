package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"log"
	"net/http"
	"strconv"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var ids = r.URL.Query()["id"]
	fmt.Println("value from productID ")
	if len(ids) > 0 {
		json.NewEncoder(w).Encode(filter(models.Products, ids))
		return
	}
	var products []models.Product
	database.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

func SaveProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	error := json.NewDecoder(r.Body).Decode(&newProduct)
	if error != nil {
		w.WriteHeader(400)
		fmt.Println("error decoding json")
	}
	result := database.DB.Create(&newProduct)
	if result.Error != nil {
		log.Println("we found a problem")
		log.Println(result.Error)
		w.WriteHeader(500)
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(result)
}

func filter(ss []models.Product, ids []string) []models.Product {
	ret := []models.Product{}
	for _, s := range ss {
		for _, i := range ids {
			convertedInt, err := strconv.Atoi(i)
			if err != nil {
				ret = nil
			}
			if s.Id == int(convertedInt) {
				ret = append(ret, s)
			}
		}
	}
	return ret
}

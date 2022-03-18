package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/models"
	"net/http"
	"strconv"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var ids = r.URL.Query()["id"]
	fmt.Println("value from productID ")
	if len(ids) > 0 {
		json.NewEncoder(w).Encode(filter(models.Products, ids))
		return
	}
	json.NewEncoder(w).Encode(models.Products)
}

func SaveProduct(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(201)
	var newProduct models.Product
	error := json.NewDecoder(r.Body).Decode(&newProduct)
	if error != nil {
		panic(error)
	}
	models.Products = append(models.Products, newProduct)
	fmt.Println(models.Products);
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

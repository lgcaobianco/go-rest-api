package routes

import (
	"go-rest-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest()  {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Hello)
	r.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	r.HandleFunc("/products", controllers.SaveProduct).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
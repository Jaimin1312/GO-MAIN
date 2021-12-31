package routes

import (
	"log"
	"net/http"
	"postgres-crud/controller"

	"github.com/gorilla/mux"
)

var router *mux.Router

func CreateRouter() {
	router = mux.NewRouter()
}

func InitializeRoutes() {
	router.HandleFunc("/customer", controller.GetAllCustomer).Methods("GET")
	router.HandleFunc("/customer", controller.CreateCustomer).Methods("POST")
	router.HandleFunc("/customer/{id}", controller.GetOneCustomer).Methods("GET")
	router.HandleFunc("/customer/{id}", controller.UpdateOneCustomer).Methods("PUT")
	router.HandleFunc("/customer/{id}", controller.DeleteOneCustomer).Methods("DELETE")
}

func StartServer() {
	log.Println("server started at 8080")
	http.ListenAndServe(":8080", router)
}

package routes

import (
	"fmt"
	"net/http"
	"restapi/configuration/controller"

	"github.com/gorilla/mux"
)

var r *mux.Router

//StartServer is started at 8082
func StartServer() {
	fmt.Println("Server is started at 8082")
	http.ListenAndServe(":8082", r)
}

//CreateRouter is...
func CreateRouter() {
	r = mux.NewRouter()
}

//InitializeRoutes is...
func InitializeRoutes() {

	r.HandleFunc("/index", controller.CustomerIndexpageProcess).Methods("GET")
	r.HandleFunc("/", controller.CustomerIndexpageProcess).Methods("GET")
	r.HandleFunc("/customer", controller.CustomerDisplaypageProcess).Methods("GET")
	r.HandleFunc("/customer", controller.CustomerInsertProcess).Methods("POST")
	r.HandleFunc("/customer/{id}", controller.CustomerEditpageProcess).Methods("GET")
	r.HandleFunc("/customer/{id}", controller.CustomerDeleteProcess).Methods("DELETE")
	r.HandleFunc("/customer/{id}", controller.CustomerUpdateProcess).Methods("PUT")
	r.HandleFunc("/error", controller.CustomerServerError).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(controller.NotFound)
}

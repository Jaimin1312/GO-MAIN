package routes

import (
	"fmt"
	"net/http"
	"packages/configuration/controller"
)

//StartServer is...
func StartServer() {
	fmt.Println("Server is started at 9092")
	http.ListenAndServe(":9092", nil)
}

//InitializeRoutes is...
func InitializeRoutes() {
	http.HandleFunc("/index", controller.CustomerIndexpageProcess)
	http.HandleFunc("/", controller.CustomerIndexpageProcess)
	http.HandleFunc("/customer", controller.CustomerDisplaypageProcess)
	http.HandleFunc("/customer-insert", controller.CustomerInsertProcess)
	http.HandleFunc("/customer-edit", controller.CustomerEditpageProcess)
	http.HandleFunc("/customer-delete", controller.CustomerDeleteProcess)
	http.HandleFunc("/customer-update", controller.CustomerUpdateProcess)
	http.HandleFunc("/customer-success", controller.CustomerSuccessProcess)
	http.HandleFunc("/customer-error", controller.CustomerServerError)

}

package main

import (
	"restapi/configuration/routes"
)

func main() {
	routes.CreateRouter()
	routes.InitializeRoutes()
	routes.StartServer()
}

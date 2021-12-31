package main

import (
	"packages/configuration/routes"
)

func main() {
	routes.InitializeRoutes()
	routes.StartServer()
}

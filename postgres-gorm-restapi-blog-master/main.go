package main

import (
	"postgres-crud/database"
	"postgres-crud/routes"
)

func main() {
	database.Initialmigration()
	routes.CreateRouter()
	routes.InitializeRoutes()
	routes.StartServer()
}

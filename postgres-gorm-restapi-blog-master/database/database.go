package database

import (
	"database/sql"
	"log"
	"postgres-crud/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *sql.DB

//connection with postgres database
func GetDatabase() *gorm.DB {
	databasename := "Customer"
	databasepassword := "1312"
	databaseurl := "postgres://postgres:" + databasepassword + "@localhost/" + databasename + "?sslmode=disable"

	connection, err := gorm.Open(postgres.Open(databaseurl), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	sqldb, err := connection.DB()
	if err != nil {
		log.Fatalln(err)
	}

	err = sqldb.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully connected to database!")
	return connection
}

//close database connetion
func Closedatabase(connection *gorm.DB) {
	sqldb, err := connection.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqldb.Close()
}

//migrate customer model to database
func Initialmigration() {
	connection := GetDatabase()
	connection.AutoMigrate(&model.Customer{})
	defer Closedatabase(connection)
	log.Println("customer model to database migration done")
}

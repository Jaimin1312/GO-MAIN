package common

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

//GetDatabase is return db connection
func GetDatabase() *sql.DB {
	db, err = sql.Open("postgres", "postgres://postgres:1312@localhost/Customer?sslmode=disable")
	CheckError(err)
	err = db.Ping()
	CheckError(err)
	fmt.Println("Database is connected")
	return db
}

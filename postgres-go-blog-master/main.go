package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//Customer is structure
type Customer struct {
	Customerid   string
	FirstName    string
	LastName     string
	Email        string
	Dateofbirth  string
	Mobilenumber string
}

//connection with postgres database
func getconnection() *sql.DB {
	databasename := "Customer"
	database := "postgres"
	databasepassword := "1312"
	databaseurl := "postgres://postgres:" + databasepassword + "@localhost/" + databasename + "?sslmode=disable"

	var err error
	db, err := sql.Open(database, databaseurl)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully connected to database!")
	return db
}

//Insert customer to database
func insertCustomer() {
	db := getconnection()
	defer db.Close()

	query := "INSERT INTO customer (firstname,lastname,email,dateofbirth,mobilenumber) VALUES($1,$2,$3,$4,$5)"
	customer := Customer{
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "Johndoe@gmail.com",
		Dateofbirth:  "13 December 1999",
		Mobilenumber: "1234567890",
	}

	_, err := db.Exec(query, customer.FirstName, customer.LastName, customer.Email, customer.Dateofbirth, customer.Mobilenumber)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Customer data inserted successfully into database")
}

//Get all customer from database
func getallCustomer() {
	db := getconnection()
	defer db.Close()

	var customers []Customer

	rows, err := db.Query("SELECT *FROM customer")
	for rows.Next() {

		var customer Customer
		err = rows.Scan(&customer.Customerid, &customer.FirstName, &customer.LastName, &customer.Email, &customer.Dateofbirth, &customer.Mobilenumber)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(customer)
		customers = append(customers, customer)
	}

}

//Update one customer to database using updateCustomerId
func updateCustomer(updateCustomerId string) {
	db := getconnection()
	defer db.Close()

	query := "UPDATE customer SET firstname = $1,lastname = $2,email = $3 ,dateofbirth =$4 ,mobilenumber =$5 where id = $6"
	customer := Customer{
		Customerid:   updateCustomerId,
		FirstName:    "John",
		LastName:     "Doe",
		Email:        "JohnDoe1312@gmail.com",
		Dateofbirth:  "13 December 1999",
		Mobilenumber: "5633234443",
	}
	_, err := db.Exec(query, customer.FirstName, customer.LastName, customer.Email, customer.Dateofbirth, customer.Mobilenumber, customer.Customerid)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Customer data updated successfully into database")
}

//Delete one customer to database using deleteCustomerId
func deleteCustomer(deleteCustomerId string) {
	db := getconnection()
	defer db.Close()

	query := "DELETE FROM customer where id = $1"
	res, err := db.Exec(query, deleteCustomerId)
	if err != nil {
		log.Fatalln(err)

	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}

	if count == 0 {
		log.Fatalln("Customer data is not deleted successfully")
	}

	log.Println("Customer data deleted successfully into database")
}

func main() {
	insertCustomer()
	getallCustomer()

	updateCustomer("1") //pass argument customerid here: 1
	getallCustomer()

	deleteCustomer("1") //pass argument customerid here: 1
	getallCustomer()
}

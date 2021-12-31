package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//Employee is struct
type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	City   string `json:"city"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}

var db *sql.DB

func checkerror(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:1312@localhost/Employee?sslmode=disable")
	checkerror(err)
	err = db.Ping()
	checkerror(err)
	fmt.Println("database is connected")
}

func getallEmployee(w http.ResponseWriter, r *http.Request) {
	query := "SELECT *FROM EMPLOYEE ORDER BY id DESC"
	rows, err := db.Query(query)
	checkerror(err)
	var employees []Employee
	for rows.Next() {
		var emp Employee
		err = rows.Scan(&emp.ID, &emp.Name, &emp.City, &emp.Mobile, &emp.Email)
		checkerror(err)
		employees = append(employees, emp)
	}
	bytedata, err := json.MarshalIndent(employees, "", "   ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytedata)
}

func saveEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee
	bodydata, err := ioutil.ReadAll(r.Body)
	checkerror(err)
	err = json.Unmarshal(bodydata, &employee)
	checkerror(err)
	query := "INSERT INTO EMPLOYEE (name,city,mobile,email) VALUES('" + employee.Name + "','" + employee.City + "','" + employee.Mobile + "','" + employee.Email + "') RETURNING id"
	fmt.Println(query)
	err = db.QueryRow(query).Scan(&employee.ID)
	checkerror(err)
	w.Header().Set("Content-Type", "application/json")
	bodydata, err = json.MarshalIndent(employee, "", "  ")
	checkerror(err)
	w.Write(bodydata)
}

func getOneEmployee(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	empid := mux.Vars(r)["id"]
	query := "SELECT *FROM EMPLOYEE WHERE id=" + empid
	rows := db.QueryRow(query)
	err := rows.Scan(&emp.ID, &emp.Name, &emp.City, &emp.Mobile, &emp.Email)
	flag := true
	if sql.ErrNoRows == err {
		flag = false
	} else {
		checkerror(err)
	}
	w.Header().Set("Content-Type", "application/json")
	bytedata, err := json.MarshalIndent(emp, "", "  ")
	checkerror(err)
	if flag {
		w.Write([]byte(bytedata))
	} else {
		w.Write([]byte("No rows found"))
	}

}

func updateOneEmployee(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	empid := mux.Vars(r)["id"]
	bodydata, err := ioutil.ReadAll(r.Body)
	checkerror(err)
	err = json.Unmarshal(bodydata, &emp)
	checkerror(err)
	query := "UPDATE EMPLOYEE SET name = $1 , city = $2 , mobile = $3 , email = $4 WHERE id = $5"
	result, err := db.Exec(query, emp.Name, emp.City, emp.Mobile, emp.Email, empid)
	_, err = result.RowsAffected()
	flag := true
	if sql.ErrNoRows == err {
		flag = false
	} else {
		checkerror(err)
	}
	w.Header().Set("Content-Type", "application/json")
	if flag {
		emp.ID, err = strconv.Atoi(empid)
		checkerror(err)
		bodydata, err = json.MarshalIndent(emp, "", " ")
		checkerror(err)
		w.Write(bodydata)
	} else {
		w.Write([]byte("No rows found"))
	}
}

func deleteOneEmployee(w http.ResponseWriter, r *http.Request) {
	empid := mux.Vars(r)["id"]
	query := "DELETE FROM EMPLOYEE WHERE id= $1"
	_, err := db.Exec(query, empid)
	checkerror(err)
	w.Write([]byte("DELETE REQUEST SUCCESS FOR " + empid))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/employees", getallEmployee).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/employees", saveEmployee).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/employees/{id}", getOneEmployee).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/employees/{id}", updateOneEmployee).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/employees/{id}", deleteOneEmployee).Methods("DELETE", "OPTIONS")
	fmt.Println("Server started at 9000")
	err := http.ListenAndServe(":9000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r))
	checkerror(err)
}

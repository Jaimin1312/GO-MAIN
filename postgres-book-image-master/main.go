package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

//Book is...
type Book struct {
	Bookid      int
	Name        string
	Image       string
	Author      string
	Price       string
	Isbn        string
	Language    string
	Description string
	CreateDate  string
}

var fm = template.FuncMap{
	"uc":    strings.ToUpper,
	"ft":    firstThree,
	"check": check,
	"inc":   increment,
}

func increment(n int) int {
	return n + 1
}

func check(n int) bool {
	return n%5 == 0
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 10 {
		s = s[:10]
	}
	return s
}

//ErrorMessage is...
type ErrorMessage struct {
	HasMessage bool
	Message    string
}

var tmpl = template.Must(template.New("").Funcs(fm).ParseGlob("template/*"))
var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:1312@localhost/Book?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Databse is connected")
}

func index(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM book ORDER BY create_date DESC")
	if err != nil {
		log.Fatalln(err)
	}

	res := []Book{}

	for rows.Next() {
		var id int
		var image []byte
		var name, author, price, isbn, createtime, language, discription string

		err = rows.Scan(&id, &name, &author, &price, &isbn, &language, &discription, &createtime, &image)

		bookimagestring := base64.StdEncoding.EncodeToString(image)
		if err != nil {
			log.Fatal(err)
		}

		res = append(res, Book{id, name, bookimagestring, author, price, isbn, language, discription, createtime})
	}
	file, _ := json.MarshalIndent(res, "", " ")
	_ = ioutil.WriteFile("Book.json", file, 0644)

	tmpl.ExecuteTemplate(w, "index.html", res)

}

func create(w http.ResponseWriter, r *http.Request) {
	errorMessage := ErrorMessage{HasMessage: false, Message: ""}
	tmpl.ExecuteTemplate(w, "create.html", errorMessage)
}

func insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		file, filenifo, err := r.FormFile("image")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			http.Redirect(w, r, "/", 301)
		}

		extension := filepath.Ext(filenifo.Filename)
		if extension != ".jpg" {
			if extension != ".jpeg" {
				if extension != ".png" {
					errorMessage := ErrorMessage{HasMessage: true, Message: "Only accept jpeg or jpg or png file"}
					tmpl.ExecuteTemplate(w, "create.html", errorMessage)
					return
				}
			}
		}

		name := r.FormValue("name")
		author := r.FormValue("author")
		price := r.FormValue("price")
		isbn := r.FormValue("isbn")
		language := r.FormValue("language")
		discription := r.FormValue("discription")
		if name == "" || author == "" || price == "" || isbn == "" || language == "" || discription == "" {
			errorMessage := ErrorMessage{HasMessage: true, Message: "Please do not Leave any empty filed"}
			tmpl.ExecuteTemplate(w, "create.html", errorMessage)
			return
		}

		createdate := time.Now()

		imageBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		sql := "INSERT INTO book(name,author,price,isbn,language,description,create_date,image) VALUES($1,$2,$3,$4,$5,$6,$7,$8)"
		_, err = db.Exec(sql, name, author, price, isbn, language, discription, createdate, imageBytes)
		if err != nil {
			panic(err.Error())
		}

		defer file.Close()
		defer fmt.Println("inserted successfully")
		http.Redirect(w, r, "/", 301)
	}

	http.Redirect(w, r, "/", 301)
}

func delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	sql := "Delete from book where bookid=$1"
	rows, err := db.Exec(sql, id)
	if err != nil {
		fmt.Println(err)
	}
	count, err := rows.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	if count == 0 {
		fmt.Println("record does not deleted")
	}
	http.Redirect(w, r, "/", 301)
}

func edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	sql := "SELECT *FROM book where bookid=" + id
	rows := db.QueryRow(sql)

	book := Book{}
	var bookstring []byte
	err := rows.Scan(&book.Bookid, &book.Name, &book.Author, &book.Price, &book.Isbn, &book.Language, &book.Description, &book.CreateDate, &bookstring)
	if err != nil {
		log.Fatalln(err)
	}
	book.Image = base64.StdEncoding.EncodeToString(bookstring)
	fmt.Println(book)
	tmpl.ExecuteTemplate(w, "edit.html", book)
}

func update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		file, filenifo, err := r.FormFile("image")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			http.Redirect(w, r, "/edit?id="+id, 301)
		}

		extension := filepath.Ext(filenifo.Filename)
		if extension != ".jpg" {
			if extension != ".jpeg" {
				if extension != ".png" {
					http.Redirect(w, r, "/edit?id="+id, 301)
					return
				}
			}
		}

		name := r.FormValue("name")
		author := r.FormValue("author")
		price := r.FormValue("price")
		isbn := r.FormValue("isbn")
		language := r.FormValue("language")
		discription := r.FormValue("discription")
		if id == "" || name == "" || author == "" || price == "" || isbn == "" || language == "" || discription == "" {
			http.Redirect(w, r, "/edit?id="+id, 301)
			return
		}

		createdate := time.Now()

		imageBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		sql := "UPDATE book SET name=$1, image=$2 ,author=$3 ,price=$4,isbn=$5,language=$6,description=$7,create_date=$8 where bookid=$9"
		rows, err := db.Exec(sql, name, imageBytes, author, price, isbn, language, discription, createdate, id)

		if err != nil {
			log.Fatalln(err)
		}
		count, err := rows.RowsAffected()
		if err != nil {
			log.Fatalln(err)

		}

		if count == 0 {
			fmt.Println("Record is not updated")
		}
		defer file.Close()
		defer fmt.Println("updated successfully")
		http.Redirect(w, r, "/", 301)
	}

	http.Redirect(w, r, "/", 301)
}

func showbook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	sql := "SELECT *FROM book where bookid=" + id
	rows := db.QueryRow(sql)

	book := Book{}
	var bookstring []byte
	err := rows.Scan(&book.Bookid, &book.Name, &book.Author, &book.Price, &book.Isbn, &book.Language, &book.Description, &book.CreateDate, &bookstring)
	if err != nil {
		log.Fatalln(err)
	}
	book.Image = base64.StdEncoding.EncodeToString(bookstring)
	tmpl.ExecuteTemplate(w, "show.html", book)
}

func main() {
	fmt.Println("Server started at 9000")
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/update", update)
	http.HandleFunc("/show", showbook)
	http.ListenAndServe(":9000", nil)
}

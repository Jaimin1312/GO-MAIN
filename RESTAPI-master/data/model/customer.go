package model

//Customer is structure
type Customer struct {
	Customerid   string `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	Dateofbirth  string `json:"dateofbirth"`
	Mobilenumber string `json:"mobilenumber"`
}

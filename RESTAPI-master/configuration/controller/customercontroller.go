package controller

import (
	"fmt"
	"net/http"
	"restapi/app/common"
	"restapi/data/model"
	"restapi/data/service"
)

var flaginsert bool
var message string
var flagupdate bool
var flagdelete bool

//CustomerIndexpageProcess is...
func CustomerIndexpageProcess(w http.ResponseWriter, r *http.Request) {
	tpl := common.GetTemplate()
	tpl.ExecuteTemplate(w, "index.html", nil)
}

//CustomerDisplaypageProcess is...
func CustomerDisplaypageProcess(w http.ResponseWriter, r *http.Request) {
	tpl := common.GetTemplate()
	customers := service.GetAllCustomer()
	var messageflag bool
	if flaginsert {
		flaginsert = false
		message = "Data Inserted successfully"
		messageflag = true
	}

	if flagupdate {
		flagupdate = false
		message = "Data Updated successfully"
		messageflag = true
	}

	if flagdelete {
		flagdelete = false
		message = "Data is Deleted successfully"
		messageflag = true
	}
	tpl.ExecuteTemplate(w, "display.html", struct {
		Customers  []model.Customer
		Message    string
		HasMessage bool
	}{customers, message, messageflag})
}

//CustomerInsertProcess is...
func CustomerInsertProcess(w http.ResponseWriter, r *http.Request) {
	flaginsert = true
	service.SaveOneCustomer(r)
	http.Redirect(w, r, "/customer", 301)
}

//CustomerEditpageProcess is ...
func CustomerEditpageProcess(w http.ResponseWriter, r *http.Request) {
	customer, err := service.GetOneCustomer(r)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/error", 301)
	}
	tpl := common.GetTemplate()
	tpl.ExecuteTemplate(w, "edit.html", customer)
}

//CustomerUpdateProcess is...
func CustomerUpdateProcess(w http.ResponseWriter, r *http.Request) {
	flagupdate = true
	service.UpdateOneCustomer(r)
}

//CustomerDeleteProcess is...
func CustomerDeleteProcess(w http.ResponseWriter, r *http.Request) {
	flagdelete = true
	service.DeleteOneCustomer(r)
}

//CustomerServerError is...
func CustomerServerError(w http.ResponseWriter, r *http.Request) {
	tpl := common.GetTemplate()
	tpl.ExecuteTemplate(w, "error.html", nil)
}

//NotFound is...
func NotFound(w http.ResponseWriter, r *http.Request) {
	tpl := common.GetTemplate()
	tpl.ExecuteTemplate(w, "404.html", nil)
}

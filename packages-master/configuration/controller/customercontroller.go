package controller

import (
	"fmt"
	"net/http"
	"packages/app/common"
	"packages/data/service"
)

//CustomerIndexpageProcess is...
func CustomerIndexpageProcess(w http.ResponseWriter, r *http.Request) {
	tpl := common.GetTemplate()
	tpl.ExecuteTemplate(w, "index.html", nil)
}

//CustomerDisplaypageProcess is...
func CustomerDisplaypageProcess(w http.ResponseWriter, r *http.Request) {
	tpl := common.GetTemplate()
	customers := service.GetAllCustomer()
	tpl.ExecuteTemplate(w, "display.html", customers)
}

//CustomerInsertProcess is...
func CustomerInsertProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		service.SaveOneCustomer(r)
		http.Redirect(w, r, "/customer-success", 301)
	}
	http.Redirect(w, r, "/", 301)
}

//CustomerEditpageProcess is ...
func CustomerEditpageProcess(w http.ResponseWriter, r *http.Request) {
	customer, err := service.GetOneCustomer(r)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/customer-error", 301)
	}
	tpl := common.GetTemplate()
	tpl.ExecuteTemplate(w, "edit.html", customer)
}

//CustomerUpdateProcess is...
func CustomerUpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		service.UpdateOneCustomer(r)
		http.Redirect(w, r, "/customer-success", 301)
	}
	http.Redirect(w, r, "/", 301)
}

//CustomerDeleteProcess is...
func CustomerDeleteProcess(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("method") == "DELETE" {
		service.DeleteOneCustomer(r)
		http.Redirect(w, r, "/customer-success", 301)
	}
	http.Redirect(w, r, "/", 301)
}

//CustomerSuccessProcess is...
func CustomerSuccessProcess(w http.ResponseWriter, r *http.Request) {
	tpl := common.GetTemplate()
	tpl.ExecuteTemplate(w, "success.html", struct{ Data string }{"Data Processed"})
}

//CustomerServerError is...
func CustomerServerError(w http.ResponseWriter, r *http.Request) {
	tpl := common.GetTemplate()
	tpl.ExecuteTemplate(w, "error.html", nil)
}

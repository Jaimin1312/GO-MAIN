package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"

	"github.com/Jeffail/gabs"
)

type ResultSuggestion struct {
	Word               []string
	SuggestionNotfound string
	HasSuggestion      bool
	Content            interface{}
	HasContent         bool
}

var tmpl = template.Must(template.ParseGlob("template/*"))
var suggetion = map[string]string{}

func getSuggestion(search string) map[string]string {

	newsearch := processquery(search)
	var url = "https://en.wikipedia.org/w/api.php?action=opensearch&search=" + newsearch

	res, _ := http.Get(url)

	links := make(map[string]string)

	cont, _ := ioutil.ReadAll(res.Body)

	result := make([][]string, 0)
	_ = json.Unmarshal(cont, &result)

	keyword, link := result[1], result[3]

	for i := 0; i < len(link); i++ {
		links[keyword[i]] = link[i]
	}

	return links
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	} else if r.Method == "POST" {

		var resultSuggestion ResultSuggestion
		search := r.FormValue("search")
		search = strings.Trim(search, " ")
		suggetion = getSuggestion(search)

		//if no suggestion found
		if len(suggetion) == 0 {

			resultSuggestion.SuggestionNotfound = search + " is not found"
			resultSuggestion.HasSuggestion = false
			resultSuggestion.Content = ""
			resultSuggestion.HasContent = false
			resultSuggestion.Word = nil
			tmpl.ExecuteTemplate(w, "index.html", resultSuggestion)
			return
		}

		words := make([]string, 0)
		for key := range suggetion {
			words = append(words, key)
		}

		resultSuggestion.SuggestionNotfound = "found"
		resultSuggestion.HasSuggestion = true
		resultSuggestion.Content = ""
		resultSuggestion.HasContent = false
		resultSuggestion.Word = words
		fmt.Println(resultSuggestion)
		tmpl.ExecuteTemplate(w, "index.html", resultSuggestion)

	}
}

func suggetionresult(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if _, ok := suggetion[query]; ok {
		query = processquery(query)
		url := "https://en.wikipedia.org/w/api.php?action=query&prop=extracts&format=json&exintro=&titles=" + query
		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error in get request ", err)
			tmpl.ExecuteTemplate(w, "index.html", nil)
			return
		}

		responsebody, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println("Error in response body ", err)
			tmpl.ExecuteTemplate(w, "index.html", nil)
			return
		}
		jsonParsed, err := gabs.ParseJSON(responsebody)
		if err != nil {
			panic(err)
		}

		Map := make(map[string]interface{})
		for key, child := range jsonParsed.Search("query", "pages").ChildrenMap() {

			fmt.Println(key)

			fmt.Println(child.Data())
			Map["extract"] = child.Data()
			fmt.Println(Map)
			Map["extract"] = Map["extract"].(map[string]interface{})["extract"]
			fmt.Println(Map)
		}

		content := Map["extract"]
		fmt.Println(content)

		var resultSuggestion ResultSuggestion
		resultSuggestion.SuggestionNotfound = ""
		resultSuggestion.HasSuggestion = false
		resultSuggestion.Content = content
		resultSuggestion.HasContent = true
		resultSuggestion.Word = nil

		fmt.Println(resultSuggestion)
		tmpl.ExecuteTemplate(w, "index.html", resultSuggestion)

	}
}

func processquery(query string) string {
	query = strings.Trim(query, " ")
	query = strings.ReplaceAll(query, " ", "_")
	return query
}

func main() {
	fmt.Println("Server started at 9000")
	http.HandleFunc("/", index)
	http.HandleFunc("/suggetionresult", suggetionresult)
	http.ListenAndServe(":9000", nil)
}

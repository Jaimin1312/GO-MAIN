package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

type Weather struct {
	Check    bool
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			fmt.Println("Error in template parse ", err)
			return
		}
		tmpl.Execute(w, nil)

	} else if r.Method == "POST" {
		tmpl, err := template.ParseFiles("index.html")

		if err != nil {
			fmt.Println("Error in template parse ", err)
			tmpl.Execute(w, nil)
			return
		}

		country := r.FormValue("country")
		fmt.Println(country)
		country = strings.Trim(country, " ")
		url := "http://api.weatherapi.com/v1/current.json?key=72d5b02f9eb4496b957105545211901&q=" + country
		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error in get request ", err)
			tmpl.Execute(w, nil)
			return
		}
		responsebody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error in converting byte ", err)
			tmpl.Execute(w, nil)
			return
		}
		var weather Weather
		err = json.Unmarshal(responsebody, &weather)
		if err != nil {
			fmt.Println("Error in unmarshal ", err)
			tmpl.Execute(w, nil)
			return
		}

		fmt.Println(weather.Error.Message)
		if weather.Error.Code == 1006 {
			weather.Check = false
			fmt.Println(weather)
			tmpl.Execute(w, weather)
		} else {
			weather.Check = true
			fmt.Println(weather)
			tmpl.Execute(w, weather)
		}
	}

}

func main() {
	fmt.Println("Server started at 8080")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

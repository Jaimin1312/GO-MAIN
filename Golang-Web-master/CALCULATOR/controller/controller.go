package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"package/model"
	"strconv"
)

func saveHistory(calculator model.Calculator) {
	historydata, err := ioutil.ReadFile("./history.json")
	if err != nil {
		fmt.Println(err)
	}
	var allcalculatorhistory []model.Calculator
	_ = json.Unmarshal(historydata, &allcalculatorhistory)

	allcalculatorhistory = append(allcalculatorhistory, calculator)
	file, err := json.MarshalIndent(allcalculatorhistory, "", " ")
	if err != nil {
		log.Fatalln("Error JSON Marshal for history file => ", err)
	}
	err = ioutil.WriteFile("history.json", file, 0644)
	if err != nil {
		log.Fatalln("Error in wrting file => ", err)
	}
}

func Operation(w http.ResponseWriter, r *http.Request) {
	var calculator model.Calculator
	json.NewDecoder(r.Body).Decode(&calculator)
	fmt.Println(calculator)
	num1, err := strconv.ParseFloat(calculator.Input1, 64)
	if err != nil {
		log.Fatal("error in string to int")
	}

	num2, err := strconv.ParseFloat(calculator.Input2, 64)
	if err != nil {
		log.Fatal("error in string to int")
	}

	calculator.Result = calculation(num1, num2, calculator.Operation)
	saveHistory(calculator)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(calculator)
}

func calculation(num1, num2 float64, operation string) string {
	var ans float64
	if operation == "+" {
		ans = float64(num1) + float64(num2)
	} else if operation == "-" {
		ans = float64(num1) - float64(num2)
	} else if operation == "*" {
		ans = float64(num1) * float64(num2)
	} else if operation == "/" {
		ans = float64(num1) / float64(num2)
	}
	return fmt.Sprint(ans)
}

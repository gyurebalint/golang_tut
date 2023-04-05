package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parseTemplate, err := template.ParseFiles("./templates/" + tmpl)
	errorCheck(err)
	err = parseTemplate.Execute(w, nil)
	errorCheck(err)
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, request *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}
func addHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello internet\n")
	sum := getSum(5, 4)
	output := fmt.Sprintf("5 + 4 = %d\n", sum)
	write(writer, output)
}

func divideHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Hello internet\n")
	sum, err := getQuotient(5, 4)
	if err != nil {
		write(writer, "cant divide by zero\n")
	} else {
		output := fmt.Sprintf("5 / 0 = %.2f", sum)
		write(writer, output)

	}
}

func getSum(x, y int) int {
	return x + y
}

func getQuotient(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	} else {
		return (x / y), nil
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/getSum", addHandler)
	http.HandleFunc("/divide", divideHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

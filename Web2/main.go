package main

import (
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
	parseTemplate, err := template.ParseFiles(("./templates/" + tmpl))
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

func main() {
	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

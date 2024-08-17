package controllers

import (
	"html/template"
	"log"
	"net/http"

	"OriD19.com/auto_complete/data/database"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Fatal("Error: Unable to parse files")
	}

	data, err := database.GetAllContacts()

	if err != nil {
		log.Fatal(err.Error())
	}

	html.Execute(w, data)
}

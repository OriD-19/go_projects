package controllers

import (
	"html/template"
	"log"
	"net/http"

	"OriD19.com/auto_complete/data/database"
)

func CreateContactGet(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/create_contact.html")

	if err != nil {
		log.Fatal(err.Error())
	}

	html.Execute(w, "Create Contact")
}

func CreateContactPost(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	phone := r.FormValue("phone")

	err := database.CreateContact(name, phone)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Could not insert the data into the database"))
		return
	}

	// Otherwise, we redirect
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}

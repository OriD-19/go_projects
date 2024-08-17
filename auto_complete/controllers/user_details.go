package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"OriD19.com/auto_complete/data/database"
)

func UserDetails(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/user_details.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	name := r.FormValue("search")
	contact, err := database.GetContact(name)

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(fmt.Errorf("Cannot find the user: %v", err.Error()).Error()))
		return
	}

	html.Execute(w, contact)
}

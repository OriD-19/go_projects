package controllers

import (
	"html/template"
	"log"
	"net/http"

	"OriD19.com/auto_complete/data/database"
)

// GetSuggestions: POST request that returns auto-complete suggestions from names
func GetSuggestions(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/suggestion.html")

	if err != nil {
		log.Fatal(err.Error())
	}

	name := r.FormValue("search")
	suggestions := database.AutoCompleteTrie.SearchPattern(name)

	html.Execute(w, suggestions)
}

package view

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"OriD19.com/api/poke/format"
	"OriD19.com/api/poke/requests"
)

func PokemonDetails(w http.ResponseWriter, r *http.Request) {
	// We have an ID parameter in the URL

	tmpl, err := template.New("details.go.html").Funcs(template.FuncMap{
		"formatName": format.FormatName,
	}).ParseFiles("templates/details.go.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Internal Server Error: %v", err.Error())))
		return
	}

	idString := r.PathValue("id")

	pokeId, err := strconv.Atoi(idString)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid pokemon id"))
		return
	}

	pokemon, err := requests.GetOne(pokeId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Internal Server Error: %v", err.Error())))
		return
	}

	tmpl.Execute(w, pokemon)

}

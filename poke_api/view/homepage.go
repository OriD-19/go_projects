package view

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"OriD19.com/api/poke/format"
	"OriD19.com/api/poke/requests"
)

func Homepage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index.go.html").Funcs(template.FuncMap{
		"formatName": format.FormatName,
	}).ParseFiles("templates/index.go.html")

	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error: Error while parsing the page"))
		return
	}

	possiblePage := r.URL.Query().Get("page")
	page, err := strconv.Atoi(possiblePage)
	fmt.Println(page)

	if err != nil {

		if possiblePage == "" {
			// If we acces the homepage, and there is no page number...
			page = 1
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error: Invalid page number"))
			return
		}
	}

	vals, err := requests.GetAllPaginated(page)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = tmpl.Execute(w, vals)

	if err != nil {
		fmt.Println(err.Error())
	}
}

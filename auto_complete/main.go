package main

import (
	"net/http"

	"OriD19.com/auto_complete/controllers"
)

func main() {
	server := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))

	server.HandleFunc("/index", controllers.HomePage)
	server.HandleFunc("POST /getSuggestions", controllers.GetSuggestions)
	server.HandleFunc("GET /createContact", controllers.CreateContactGet)
	server.HandleFunc("POST /createContact", controllers.CreateContactPost)
	server.HandleFunc("GET /userDetails", controllers.UserDetails)
	server.Handle("/", fs)

	http.ListenAndServe(":3306", server)
}

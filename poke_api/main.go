package main

import (
	"fmt"
	"net/http"

	"OriD19.com/api/poke/view"
)

const port = ":3306"

func main() {

	server := http.NewServeMux()

	server.HandleFunc("/", view.Homepage)
	server.HandleFunc("GET /pokemon/{id}", view.PokemonDetails)

	fs := http.FileServer(http.Dir("./templates"))
	server.Handle("/static", fs)

	fmt.Println("Server about to start on port", port)
	http.ListenAndServe(port, server)
}

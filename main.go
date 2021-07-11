package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Foobar struct {
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type Foobars []Foobar

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>Hello world GO</b>")
}

func allFoobars(w http.ResponseWriter, r *http.Request) {
	foobars :=
		Foobars{
			Foobar{Name: "Name test", Description: "Description test"},
		}
	fmt.Println("All foobars")
	json.NewEncoder(w).Encode(foobars)
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/json", allFoobars)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
} 
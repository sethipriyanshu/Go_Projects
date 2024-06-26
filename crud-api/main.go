package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//"encoding/json"  // help encode data
	//"math/rand"
	//"net/http"
	//"strconv"
	//"github.com/gorilla/mux"
)

// declare a struct of type movie
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// declare a struct of type director
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// declare a slice to store the movies (not using a backend database for now)
var movies []Movie

func main() {
	// declare a new Route
	r := mux.NewRouter()

	// handling functions on respective routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("POST")

	fmt.Printf("Starting Server at Port 8000\n")
	//
	log.Fatal(http.ListenAndServe(":8000", r))
}

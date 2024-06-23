package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// initialize a file server using the http import and set it to static directory
	fileServer := http.FileServer(http.Dir("./static"))
	// setup handling for files and function routes using handler functions
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// display the start of server
	fmt.Printf("Starting server at port 8080 \n")
	// error handling and creating the server using ListenAndServe function
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// using log package to log errors
		log.Fatal(err)
	}
}

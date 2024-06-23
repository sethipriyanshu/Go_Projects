package main

import (
	"fmt"
	"log"
	"net/http"
)

// every api/route has a response or a request and that is what the parameters we pass signify * is a pointer to the request user sends
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// handling the situation if the request does not come from hello route
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// we dont want people to post to /hello but only be able to get
	if r.Method != "GET" {
		http.Error(w, "Method is not Supported", http.StatusNotFound)
		return
	}
	// Fprintf function formats and writes to w
	fmt.Fprintf(w, "Hello!")

}

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

package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/test" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "You can't use any other method then GET", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Golang webserver")
}

func main() {
	http.HandleFunc("/test", helloHandler) // Update this line of code

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

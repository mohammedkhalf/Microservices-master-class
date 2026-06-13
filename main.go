package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Go server on port 9090!")
}

func main() {
	http.HandleFunc("/", home)

	http.HandleFunc("/goodbye/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Good bye world")
	})

	fmt.Println("Server running on http://localhost:9090")

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

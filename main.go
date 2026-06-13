package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func home(rw http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Oooop"))
		return
	}

	println(string(body))
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

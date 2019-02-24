package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Go API!!")
	})

	port := ":8080"

	fmt.Println("start!", "http://localhost"+port)
	log.Fatal(http.ListenAndServe(port, nil))
}

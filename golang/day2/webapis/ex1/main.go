package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Go docker microcontainer")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello test")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

// FunctionToTest testing
func FunctionToTest() string {
	return "my test"
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var clients = []string{"1", "2", "3"}

func main() {
	fmt.Println("Go docker microcontainer")

	router := mux.NewRouter()
	router.HandleFunc("/client", getClients).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func getClients(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(clients)
}

// FunctionToTest testing
func FunctionToTest() string {
	return "my test"
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var clients = []string{"1", "2", "3", "4", "5", "6"}

func main() {
	fmt.Println("Go docker microcontainer")

	router := mux.NewRouter()
	router.HandleFunc("/client", getClients).Methods("GET", "POST")
	router.HandleFunc("/client/{id}", getClients).Methods("GET")
	//router.HandleFunc("/client", createClient).Methods("POST")
	router.HandleFunc("/client", updateClient).Methods("PUT")
	router.HandleFunc("/client", deleteClient).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func getClients(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(clients)
}

func createClient(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("posted")
}

func updateClient(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("updated")
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("deleted")
}

// FunctionToTest testing
func FunctionToTest() string {
	return "my test"
}

package main

import (
	"app/db"
	"app/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Go docker microcontainer")

	router := mux.NewRouter()
	router.HandleFunc("/client", getClients).Methods("GET")
	router.HandleFunc("/client/{id}", getClient).Methods("GET")
	router.HandleFunc("/client", createClient).Methods("POST")
	router.HandleFunc("/client", updateClient).Methods("PUT")
	router.HandleFunc("/client/{id}", deleteClient).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func getClients(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.Client)
}
func getClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, _ := strconv.Atoi(params["id"])
	for _, client := range db.Client {
		if client.ID == i {
			json.NewEncoder(w).Encode(client)
			break
		}
	}
}

func createClient(w http.ResponseWriter, r *http.Request) {
	var client model.Client
	_ = json.NewDecoder(r.Body).Decode(&client)
	db.Client = append(db.Client, client)
	json.NewEncoder(w).Encode(db.Client)
}

func updateClient(w http.ResponseWriter, r *http.Request) {
	var client model.Client
	_ = json.NewDecoder(r.Body).Decode(&client)
	for i, value := range db.Client {
		if value.ID == client.ID {
			db.Client[i] = client
		}
	}
	json.NewEncoder(w).Encode(db.Client)
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	for i, client := range db.Client {
		if client.ID == id {
			// nasty hack, we make a copy of an array skiping the (delete) paramater id value ( if exist )
			db.Client = append(db.Client[:i], db.Client[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(db.Client)
}

// FunctionToTest testing
func FunctionToTest() string {
	return "my test"
}

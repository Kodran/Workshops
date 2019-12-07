package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	api "github.com/kodran/workshop/day3/webapi/api/client"
)

func main() {
	fmt.Println("Go docker microcontainer")

	router := mux.NewRouter()
	router.HandleFunc("/client", api.GetClients).Methods("GET")
	router.HandleFunc("/client/{id}", api.GetClient).Methods("GET")
	router.HandleFunc("/client", api.CreateClient).Methods("POST")
	router.HandleFunc("/client", api.UpdateClient).Methods("PUT")
	router.HandleFunc("/client/{id}", api.DeleteClient).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}

// FunctionToTest testing
func FunctionToTest() string {
	return "my test"
}

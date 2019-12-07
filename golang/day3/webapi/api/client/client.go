package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kodran/workshop/day3/webapi/db"
	"github.com/kodran/workshop/day3/webapi/model"

	"github.com/gorilla/mux"
)

func GetClients(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.Client)
}
func GetClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, client := range db.Client {
		if client.ID == id {
			json.NewEncoder(w).Encode(client)
			break
		}
	}
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	var client model.Client
	_ = json.NewDecoder(r.Body).Decode(&client)
	db.Client = append(db.Client, client)
	json.NewEncoder(w).Encode(db.Client)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	var client model.Client
	_ = json.NewDecoder(r.Body).Decode(&client)
	for i, value := range db.Client {
		if value.ID == client.ID {
			db.Client[i] = client
		}
	}
	json.NewEncoder(w).Encode(db.Client)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])
	for i, client := range db.Client {
		if client.ID == id {
			db.Client = append(db.Client[:i], db.Client[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(db.Client)
}

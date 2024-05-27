package controller

import (
	"encoding/json"
	"fmt"
	"goserver/config"
	"goserver/model"
	store "goserver/sqlstore"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	db = config.GetDb()
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	params := mux.Vars(r)
	err := store.GetUser(&user, db, params["id"])
	if err != nil {
		http.Error(w, "Failed to get Users", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user []model.User
	err := store.GetUsers(&user, db)
	if err != nil {
		http.Error(w, "Failed to get Users", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	json.NewDecoder(r.Body).Decode(&user)
	err := store.CreateUser(&user, db)
	if err != nil {
		fmt.Println("Failed to create new data")
		http.Error(w, "Failed to create new data", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	params := mux.Vars(r)
	err := store.DeleteUser(&user, db, params["id"])
	if err != nil {
		http.Error(w, "Failed to delete", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode("User deletion seccessful")
	json.NewEncoder(w).Encode(params["id"])

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser model.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Failed to decode JSON request body", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	userID := params["id"]
	var existingUser model.User
	if err := store.GetUser(&existingUser, db, userID); err != nil {
		http.Error(w, "Failed to get user from the database", http.StatusNotFound)
		return
	}

	existingUser.Name = newUser.Name
	existingUser.Email = newUser.Email
	existingUser.Age = newUser.Age

	if err := store.UpdateUser(&existingUser, db); err != nil {
		http.Error(w, "Failed to update user in the database", http.StatusInternalServerError)
		return
	}

	// Encode the updated user as JSON and send it in the response
	if err := json.NewEncoder(w).Encode(existingUser); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}

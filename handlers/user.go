package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nigerdyanes/gorm-crud/models"
	"gorm.io/gorm"
)

func GetUsers(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
	var users = []models.User{}
	db.Find(&users)
	output, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(output))
}

func GetUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
	var user = models.User{}
	vars := mux.Vars(r)
	db.Find(&user, "id = ?", vars["id"])
	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(output))
}

func CreateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
	var user = models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	db.Create(&user)
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(output))
}

func UpdateUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
	var user = models.User{}
	vars := mux.Vars(r)
	db.Find(&user, "id = ?", vars["id"])
	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	db.Save(&user)
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(output))
}

func DeleteUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Header().Set("Content-Type", "application/json")
	var user = models.User{}
	vars := mux.Vars(r)
	db.Find(&user, "id = ?", vars["id"])
	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	db.Delete(&user)
	fmt.Fprintln(w, "User deleted")
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nigerdyanes/gorm-crud/database"
	"github.com/nigerdyanes/gorm-crud/handlers"
	"github.com/nigerdyanes/gorm-crud/models"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUsers(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/api/v1/users/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetUser(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUser(w, r, db)
	}).
		Methods("POST")

	r.HandleFunc("/api/v1/users/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateUser(w, r, db)
	}).
		Methods("PUT")

	r.HandleFunc("/api/v1/users/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteUser(w, r, db)
	}).
		Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

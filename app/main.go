package main

import (
	"github.com/gorilla/mux"
	"github.com/t0masphp/golang-todo/controllers"
	"github.com/t0masphp/golang-todo/database"
	"github.com/t0masphp/golang-todo/models"
	"net/http"
)

func main() {
	db := database.Init()
	db.AutoMigrate(&models.TodoItemModel{})
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.GetEntries).Methods("GET")
	router.HandleFunc("/", controllers.CreateEntry).Methods("POST")
	err := http.ListenAndServe(":8888", router)
	if err != nil {
		return
	}
}

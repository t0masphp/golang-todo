package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/t0masphp/golang-todo/database"
	"github.com/t0masphp/golang-todo/models"
	"io"
	"net/http"
)

func GetEntries(writer http.ResponseWriter, request *http.Request) {
	var todos []models.TodoItemModel
	db := database.Init()
	db.Find(&todos)
	data, _ := json.Marshal(todos)
	writer.Header().Set("Content-Type", "application/json")
	_, err := io.WriteString(writer, string(data))
	if err != nil {
		return
	}
}

func CreateEntry(writer http.ResponseWriter, request *http.Request) {
	db := database.Init()
	var item models.TodoItemModel
	_ = json.NewDecoder(request.Body).Decode(&item)
	fmt.Println(item, request.Body)
	db.Create(item)
	data, _ := json.Marshal(item)
	writer.Header().Set("Content-Type", "application/json")
	_, err := io.WriteString(writer, string(data))
	if err != nil {
		return
	}
}

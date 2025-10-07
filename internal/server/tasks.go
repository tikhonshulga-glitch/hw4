package server

import (
	"encoding/json"
	"hwproject/internal/domain/tasks/models"
	"hwproject/internal/service/taskservice"
	"net/http"
	"strconv"
	"strings"
)

func (srv *RentAPI) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks := taskservice.GetInMemoryTasks()

	json.NewEncoder(w).Encode(tasks)

}

func (srv *RentAPI) CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTask, err := taskservice.CreateInMemoryTask(newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}

func (srv *RentAPI) GetTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := taskservice.GetInMemoryTask(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (srv *RentAPI) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/update/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask models.Task

	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTask.ID = id

	updatedTaskPtr, err := taskservice.UpdateInMemoryTask(updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTaskPtr)
}

func (srv *RentAPI) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/delete/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	deletedTask, err := taskservice.DeleteInMemoryTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deletedTask)
}

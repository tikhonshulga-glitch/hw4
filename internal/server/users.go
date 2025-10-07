package server

import (
	"encoding/json"
	"hwproject/internal/domain/users/models"
	"hwproject/internal/service/userservice"
	"net/http"
	"strconv"
	"strings"
)

func (srv *RentAPI) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := userservice.GetInMemoryUsers()

	json.NewEncoder(w).Encode(users)
}

func (srv *RentAPI) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := userservice.CreateInMemoryUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)

}

func (srv *RentAPI) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := userservice.GetInMemoryUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (srv *RentAPI) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/update/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusNotFound)
	}

	var updatedUser models.User

	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser.ID = id

	updatedUserPtr, err := userservice.UpdateInMemoryUser(updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUserPtr)

}

func (srv *RentAPI) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/delete/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	deletedUser, err := userservice.DeleteInMemoryUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deletedUser)

}

func (srv *RentAPI) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req models.UserReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	token, err := userservice.LoginInMemoryUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

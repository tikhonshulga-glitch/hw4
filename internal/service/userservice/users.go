package userservice

import (
	"fmt"
	"hwproject/internal/domain/users/models"
	"hwproject/internal/repository/inmemory"

	"golang.org/x/crypto/bcrypt"
)

func GetInMemoryUsers() []models.User {
	return inmemory.Users
}

func CreateInMemoryUser(newData models.User) (models.User, error) {
	for _, v := range inmemory.Users {
		if newData.Email == v.Email {
			return models.User{}, fmt.Errorf("user with email %s already exists", newData.Email)
		}
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newData.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to hash password: %v", err)
	}

	newData.Password = string(hash)
	newData.ID = len(inmemory.Users) + 1

	inmemory.Users = append(inmemory.Users, newData)

	return newData, nil
}

func GetInMemoryUser(userID int) (*models.User, error) {
	for i := range inmemory.Users {
		if userID == inmemory.Users[i].ID {
			return &inmemory.Users[i], nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", userID)
}

func UpdateInMemoryUser(newData models.User) (*models.User, error) {
	for i, v := range inmemory.Users {
		if v.ID == newData.ID {
			inmemory.Users[i] = newData
			return &inmemory.Users[i], nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", newData.ID)
}

func DeleteInMemoryUser(userID int) (*models.User, error) {
	for i := range inmemory.Users {
		if inmemory.Users[i].ID == userID {
			deletedTask := inmemory.Users[i]
			inmemory.Users = append(inmemory.Users[:i], inmemory.Users[i+1:]...)
			return &deletedTask, nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", userID)
}

func LoginInMemoryUser(newData models.UserReq) (string, error) {
	for i := range inmemory.Users {
		v := &inmemory.Users[i]
		if newData.Email == v.Email {

			if err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(newData.Password)); err != nil {
				return "", fmt.Errorf("incorrect login or password")
			}
			return "registered", nil
		}
	}
	return "", fmt.Errorf("incorrect login or password")
}

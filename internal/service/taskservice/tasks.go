package taskservice

import (
	"fmt"
	"hwproject/internal/domain/tasks/models"
	"hwproject/internal/repository/inmemory"
)

func GetInMemoryTasks() []models.Task {
	return inmemory.Tasks
}

func CreateInMemoryTask(t models.Task) (models.Task, error) {

	t.ID = len(inmemory.Tasks) + 1

	if t.Tittle == "" {
		t.Tittle = "Untitled"
	}

	maxId := 0
	for _, task := range inmemory.Tasks {
		if task.ID > maxId {
			maxId = task.ID
		}
	}

	if t.Status != "New" && t.Status != "In Process" && t.Status != "Done" {
		return models.Task{}, fmt.Errorf("task with ID %d has incorrect status", t.ID)
	}

	inmemory.Tasks = append(inmemory.Tasks, t)

	return t, nil
}

func GetInMemoryTask(taskID int) (*models.Task, error) {
	for i := range inmemory.Tasks {
		if taskID == inmemory.Tasks[i].ID {
			return &inmemory.Tasks[i], nil
		}
	}
	return nil, fmt.Errorf("task with ID %d not found", taskID)
}

func UpdateInMemoryTask(newData models.Task) (*models.Task, error) {

	for i, v := range inmemory.Tasks {
		if v.ID == newData.ID {
			if newData.Status != "New" && newData.Status != "In process" && newData.Status != "Done" {
				return nil, fmt.Errorf("task with ID %d has incorrect status", newData.ID)
			}
			inmemory.Tasks[i] = newData
			return &inmemory.Tasks[i], nil
		}
	}

	return nil, fmt.Errorf("task with ID %d not found", newData.ID)
}

func DeleteInMemoryTask(taskID int) (*models.Task, error) {
	for i := range inmemory.Tasks {
		if inmemory.Tasks[i].ID == taskID {
			deletedTask := inmemory.Tasks[i]
			inmemory.Tasks = append(inmemory.Tasks[:i], inmemory.Tasks[i+1:]...)
			return &deletedTask, nil
		}
	}

	return nil, fmt.Errorf("task with ID %d not found", taskID)
}

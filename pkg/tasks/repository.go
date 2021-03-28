package tasks

import (
	"github.com/jkbmdk/kanban-api/internal/database"
	"github.com/jkbmdk/kanban-api/internal/models"
)

func GetAllTasks() []models.Task {
	var tasks []models.Task
	database.DB.Find(&tasks)
	return tasks
}

func GetTaskByID(id int) (models.Task, error) {
	var task models.Task
	err := database.DB.Where("id = ?", id).First(&task).Error
	return task, err
}

func StoreTask(task *models.Task) error {
	err := database.DB.Create(task).Error
	return err
}

func UpdateTask(task *models.Task) error {
	err := database.DB.Save(task).Error
	return err
}

func DeleteTask(id int) error {
	task, err := GetTaskByID(id)
	if err != nil {
		return err
	}
	return database.DB.Delete(task).Error
}
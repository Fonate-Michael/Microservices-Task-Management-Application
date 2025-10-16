package controllers

import (
	"net/http"
	"task/db"
	"task/model"

	"github.com/gin-gonic/gin"
)

func GetTask(context *gin.Context) {
	row, err := db.DB.Query("SELECT * FROM tasks")

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch Tasks ma boi"})
		return
	}

	var tasks []model.Task
	for row.Next() {
		var task model.Task
		if err := row.Scan(&task.Id, &task.User_id, &task.Title, &task.Description); err != nil {
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch Tasks ma boi"})
			return
		}
		tasks = append(tasks, task)
	}
	context.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

func AddTask(context *gin.Context) {
	var newTask model.Task
	userId := context.MustGet("user_id")

	err := context.BindJSON(&newTask)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to bind json"})
		return

	}

	_, err = db.DB.Exec("INSERT INTO tasks(user_id, title, description) VALUES($1, $2, $3)", userId, newTask.Title, newTask.Description)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to add task ma boi"})
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Tasks Created successfully!"})
}

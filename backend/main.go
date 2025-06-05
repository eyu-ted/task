package main

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var tasks = []Task{
	{ID: 1, Title: "asbeza megzat", Completed: false},
	{ID: 2, Title: "bet mewelwel", Completed: true},
}

func get(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}
func create(c *gin.Context) {
	var newTask Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if strings.TrimSpace(newTask.Title) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title cannot be empty"})
		return
	}
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, newTask)
}
func update(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}
func delete(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
}
func main() {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/tasks", get)
		api.POST("/tasks", create)
		api.PUT("/tasks/:id", update)
		api.DELETE("/tasks/:id", delete)
	}
	r.Run(":8080")
}

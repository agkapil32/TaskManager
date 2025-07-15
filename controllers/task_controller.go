package controllers

import (
	"TaskManager/models"
	"TaskManager/services"
	"TaskManager/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Invalid input")
		return
	}
	newTask := services.CreateTask(task)
	c.JSON(http.StatusCreated, newTask)
}

func GetTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := services.GetTaskByID(id)
	if err != nil {
		utils.JSONError(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Invalid input")
		return
	}
	updated, err := services.UpdateTask(id, task)
	if err != nil {
		utils.JSONError(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, updated)
}

func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := services.DeleteTask(id)
	if err != nil {
		utils.JSONError(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func ListTasks(c *gin.Context) {
	filters := map[string]string{
		"userId": c.Query("userId"),
		"status": c.Query("status"),
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	sortBy := c.DefaultQuery("sort", "id")

	tasks := services.ListTasks(filters, page, pageSize, sortBy)
	c.JSON(http.StatusOK, tasks)
}

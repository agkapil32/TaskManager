package routes

import (
	"TaskManager/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	task := router.Group("/tasks")
	{
		task.POST("", controllers.CreateTask)
		task.GET("", controllers.ListTasks)
		task.GET("/:id", controllers.GetTask)
		task.PUT("/:id", controllers.UpdateTask)
		task.DELETE("/:id", controllers.DeleteTask)
	}
}

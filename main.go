package main

import (
	"TaskManager/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	routes.RegisterRoutes(r)
	log.Println("TaskManager service started successfully on port 8080")
	r.Run(":8080")
}

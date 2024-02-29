package main

import (
	"github.com/rezaakbar35/final-task/controller/userController"

	"github.com/rezaakbar35/final-task/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDB()

	r.GET("/api/users", userController.GetAllUser)
	r.GET("/api/users/:id", userController.GetUserById)
	r.POST("/api/users", userController.CreateUser)
	r.PUT("/api/users/:id", userController.UpdateUser)
	r.DELETE("/api/users/:id", userController.DeleteUser)

	r.Run()
}
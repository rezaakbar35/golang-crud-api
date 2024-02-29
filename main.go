package main

import (
	"github.com/rezaakbar35/golang-crud-api/controller/userController"
	"github.com/rezaakbar35/golang-crud-api/controller/photoController"
	"github.com/rezaakbar35/golang-crud-api/model"
	"github.com/rezaakbar35/golang-crud-api/database"
    "github.com/rezaakbar35/golang-crud-api/database/migration"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := database.ConnectDB()
    database.RunMigrations(db)
    defer db.Close()

	r.GET("/api/users", userController.GetAllUser)
	r.GET("/api/users/:id", userController.GetUserById)
	r.POST("/api/users", userController.CreateUser)
	r.PUT("/api/users/:id", userController.UpdateUser)
	r.DELETE("/api/users/:id", userController.DeleteUser)

	r.GET("/api/photos", photoController.GetAllPhoto)
	r.GET("/api/photos/:id", photoController.GetPhotoById)
	r.POST("/api/photos", photoController.CreatePhoto)
	r.PUT("/api/photos/:id", photoController.UpdatePhoto)
	r.DELETE("/api/photos/:id", photoController.DeletePhoto)

	r.Run()
}
package migration

import (
    "gorm.io/gorm"
    "github.com/rezaakbar35/golang-crud-api/model"
    "github.com/rezaakbar35/golang-crud-api/database"
)

func Migrate() {
    db := database.ConnectDB()
    defer db.Close()

    // Add your migration logic here
    db.AutoMigrate(&model.User{}, &model.Photo{})
}

func Rollback() {
    db := database.ConnectDB()
    defer db.Close()

    // Add your rollback logic here
    db.Migrator().DropTable(&model.User{}, &model.Photo{})
}

package migration

import (
    "github.com/rezaakbar35/golang-crud-api/database"
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
    // Add your migration logic here
    db.AutoMigrate(&model.User{}, &model.Photo{})
}

func Rollback(db *gorm.DB) {
    // Add your rollback logic here
    db.Migrator().DropTable(&model.User{}, &model.Photo{})
}

package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "github.com/rezaakbar35/golang-crud-api/model"
    "github.com/rezaakbar35/golang-crud-api/database/migration"
)

func ConnectDB() *gorm.DB {
    dsn := "host=localhost user=postgres password=re0800za dbname=golang-final-task port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    return db
}

func RunMigrations(db *gorm.DB) {
    migration.Migrate(db)
}

func RollbackMigrations(db *gorm.DB) {
    migration.Rollback(db)
}

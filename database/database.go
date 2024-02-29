package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
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
    // No need to pass `db` to `migration.Migrate` function
    // It's already connected
    migration.Migrate()
}

func RollbackMigrations(db *gorm.DB) {
    // No need to pass `db` to `migration.Rollback` function
    // It's already connected
    migration.Rollback()
}

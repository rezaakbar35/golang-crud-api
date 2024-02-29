package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id 			uint 		`gorm:"primaryKey" json:"id"`
	Username 	string 		`json:"username"`
	Email 		string 		`json:"email"`
	Password 	string 		`json:"password"`
	CreatedAt 	time.Time 	`json:"createdAt"`
	UpdatedAt 	time.Time 	`json:"updatedAt"`
	Photos 		[]Photo 	`json:"photos"`
}
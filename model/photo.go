package model

import (
	"gorm.io/gorm"
)

type Photo struct{
	gorm.Model
	Id 			uint 		`gorm:"primaryKey" json:"id"`
	Title 		string 		`json:"photoUrl"`
	Caption 	string 	 	`json:"createdAt"`
	photoUrl 	string	 	`json:"updatedAt"`
	UserId 		uint 		`json:"userId"`
	User 		User 		`json:"user"`
}
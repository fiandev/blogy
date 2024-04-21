package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title 		string `gorm:"not null"`
	Slug    	string `gorm:"uniqueIndex;not null"`
	Content  	string `gorm:"not null"`
	AuthorId  	uint8
}
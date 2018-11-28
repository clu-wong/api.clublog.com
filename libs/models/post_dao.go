package models

import "github.com/jinzhu/gorm"

type Post struct{
	gorm.Model
	Title string `gorm:"type:varchar(100);unique_index"`
	Content string `gorm:"type:text"`
	Author User
}


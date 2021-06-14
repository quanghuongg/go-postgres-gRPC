package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `form:"Name" json:"Name" xml:"Name"  binding:"required"`
	Id int
}

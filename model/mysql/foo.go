package mysql

import "github.com/jinzhu/gorm"

type Foo struct {
	gorm.Model
	Id       int    `gorm:"primary_key:unique"`
	Name     string `gorm:"no null"`
	Password string `gorm:"no null"`
}

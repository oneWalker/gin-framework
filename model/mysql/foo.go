package mysql

import (
	mysql "gin-practice/pkg/db/mysql"
	"github.com/jinzhu/gorm"
)

type Foo struct {
	*gorm.Model
	Id       int64  `gorm:"primary_key:unique"`
	Name     string `gorm:"no null"`
	Password string `gorm:"no null"`
}

var Foos []Foo

func (f *Foo) Insert(db *gorm.DB) (id int64, err error) {
	create := db.Create(&f)
	id = f.Id
	if create.Error != nil {
		return 0, create.Error
	}

	return id, nil
}

func (f *Foo) Get() (foos []Foo, err error) {
	if err = mysql.DB.Find(&foos).Error; err != nil {
		return
	}
	return
}

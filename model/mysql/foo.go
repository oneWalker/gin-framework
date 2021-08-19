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

func (f *Foo) GetAll() (foos []Foo, err error) {
	//默认会有deleted_at的条件判断，所以先直接写死
	if err = mysql.DB.Where("deleted_at is null").Unscoped().Find(&foos).Error; err != nil {
		return
	}
	return
}

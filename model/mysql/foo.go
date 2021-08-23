package mysql

import (
	"fmt"
	mysql "gin-practice/pkg/db/mysql"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type Foo struct {
	*gorm.Model
	// 只有注释掉才能得到对应的值
	// ID       int64  `gorm:"primary_key:unique"`
	Name     string `gorm:"no null"`
	Password string `gorm:"no null"`
}

var Foos []Foo

func (f *Foo) Insert(foo *Foo) (id int, err error) {
	create := mysql.DB.Create(&foo)
	fmt.Println(create.Value, foo.ID)
	id = int(foo.ID)
	if create.Error != nil {
		return 0, create.Error
	}
	return id, nil
}

func (f *Foo) GetAll() (foos []Foo, err error) {
	// //默认会有deleted_at的条件判断，所以使用忽略条件Where("deleted_at is null").Unscoped()
	// result :=mysql.DB.Where("deleted_at is null").Unscoped().Find(&foos);

	//条件查询获取列表
	//result := mysql.DB.Where("deleted_at is null").Unscoped().Find(&foos, []int{5, 11, 15})

	//String条件
	//result := mysql.DB.Where("name = ? AND password >= ?", "test2", "test2").Find(&foos)
	//Struct & Map条件
	//result := mysql.DB.Where(&Foo{Name: "test2", Password: "test2"}).Find(&foos)
	result := mysql.DB.Where(map[string]interface{}{"name": "test2", "password": "test2"}).Offset(3).Limit(2).Order("ID desc").Find(&foos)

	fmt.Println(result)
	//对于结果进行循环
	for k, v := range foos {
		v.Name = "name"
		foos[k] = v
		fmt.Println("foo", k, v)
	}
	//利用主键进行切片
	//result := mysql.DB.Where([]int{5, 11, 15}).Find(&foos)
	if result.Error != nil {
		return
	}
	return
}

func (f *Foo) GetOne() (foo Foo, err error) {

	//获取第一条记录，不加id条件的时候则不会指定id
	result := mysql.DB.Where("deleted_at is null").Unscoped().First(&foo, 5)

	// //获取一条记录，没有指定排序顺序
	// result := mysql.DB.Take(&foo, 5)

	//获取最后一条记录(主键降序)
	//result := mysql.DB.Last(&foo)

	if result.Error != nil {
		logrus.Fatalf("select error %v", result.Error)
	}
	return
}

func (f *Foo) Update(foo *Foo) (err error) {
	//事物和相应的更新函数
	err = nil
	tx := mysql.DB.Begin()
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
		}
	}()
	//不使用Hooks的话则定义使用UpdateColumn，UpdateColumns
	tx.Model(&Foo{}).Where("id = ?", 80).Updates(&foo)
	//tx.Model(&Foo{}).Where("active = ?", true).Where("id = ?", 80).Update("name", "rename")
	//tx.Rollback()
	tx.Commit()
	return err
}

func (f *Foo) Delete(id int) (err error) {
	//根据主键删除
	//result := mysql.DB.Delete(&Foo{},10)
	//批量删除，没有任何条件的情况下执行批量删除，GORM 不会执行该操作，并返回 ErrMissingWhereClause 错误对此，你必须加一些条件，或者使用原生 SQL，或者启用 AllowGlobalUpdate
	err = mysql.DB.Where("id = ?", id).Delete(&Foo{}).Error
	//original sql dialect
	//mysql.DB.Exec("DELETE FROM foos")

	//result := mysql.DB.Delete(Foo{},"name LIKE ?","%test2%")
	return err
}

//不实现功能，只是作为样例参考，即没有经过测试
func (f *Foo) Demeo() {
	// 迁移 schema
	mysql.DB.AutoMigrate(&Foo{})

	// Create
	mysql.DB.Create(&Foo{Name: "liukun", Password: "liukun"})

	// Read
	var foo Foo
	mysql.DB.First(&foo, 5)             // 根据整形主键查找
	mysql.DB.First(&foo, "id = ?", "5") // 查找 code 字段值为 D42 的记录

	// Update - 将查找到的去进行更新
	mysql.DB.Model(&foo).Update("Price", 200)
	// Update - 更新多个字段
	mysql.DB.Model(&foo).Updates(Foo{Name: "updatetest", Password: "updatetest"}) // 仅更新非零值字段
	mysql.DB.Model(&foo).Updates(map[string]interface{}{"Name": "updatetest", "Password": "updatetest"})

	// Delete - 删除foo当中的id
	mysql.DB.Delete(&foo, 5)
}

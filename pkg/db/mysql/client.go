package mysql

import (
	config "gin-practice/config"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	DB *gorm.DB
)

//Init初始化数据库
func Init() (err error) {
	env := os.Getenv("ENV")
	var dsn string
	switch env {
	case "pro":
		dsn = config.ProDsn()
	case "test":
		dsn = config.TestDsn()
	default:
		dsn = config.DevDsn()
	}

	if DB, err = gorm.Open("mysql", dsn); err != nil {
		logrus.Fatalf("mysql connect failed: %v", err)
	}
	DB.LogMode(true)

	if err = DB.DB().Ping(); err != nil {
		log.Fatalf("database heartbeat failed: %v", err)
	}

	logrus.Info("mysql connect successfully")
	return err
}

//close Mysql
func Close() error {
	if DB != nil {
		if err := DB.Close(); err != nil {
			return err
		}
	}

	logrus.Println("mysql connect closed")
	return nil
}

package mongodb

import (
	"context"
	"gin-practice/config"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Mongodb *mongo.Database
	client  *mongo.Client
)

//mongodb数据库进行初始化
func Init() error {
	var err error
	env := os.Getenv("ENV")
	var uri string
	switch env {
	case "test":
		uri = config.TestMongo()
	case "pro":
		uri = config.ProMongo()
	default:
		uri = config.DevMongo()
	}
	option := options.Client().ApplyURI(uri)
	//设置连接池个数
	option.SetMaxPoolSize(10)
	//设置连接超时时间
	option.SetConnectTimeout(5 * time.Second)

	//1.建立连接
	if client, err = mongo.Connect(context.TODO(), option); err != nil {
		logrus.Fatalf("couldn't connect to mongo: %v", err)
		return err
	}
	//选择数据库
	Mongodb = client.Database("test")
	logrus.Info("mongo connect successfully")
	return nil
}

//mongodb数据库进行关闭
func Close() error {
	if client != nil {
		if err := client.Disconnect(nil); err != nil {
			return err
		}
	}
	logrus.Info("mongo connect disconnected")
	return nil
}

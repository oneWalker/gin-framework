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
	mongodb *mongo.Database
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
	//1.建立连接
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetConnectTimeout(5*time.Second)); err != nil {
		logrus.Fatalf("couldn't connect to mongo: %v", err)
		return err
	}
	//选择数据库
	mongodb = client.Database("test")
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

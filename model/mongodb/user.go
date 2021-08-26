package mongodb

import (
	"context"
	"gin-practice/pkg/db/mongodb"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type User struct {
	Username string
	Password string
	Name     string
	Gender   string
}

var (
	Users []User
)

func (u *User) InsertOne(user *User) (id string, err error) {
	collection := mongodb.Mongodb.Collection("users")
	//插入一个
	insertResult, err := collection.InsertOne(context.TODO(), &user)
	//插入多个
	//insertResults, err := collection.InsertOne(context.TODO(), &Users)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("type%v%v", insertResult.InsertedID, reflect.TypeOf(insertResult.InsertedID))
	mongoId := insertResult.InsertedID
	id = mongoId.(primitive.ObjectID).Hex()
	return id, err
}

func (u *User) GetAll() (users *[]User, err error) {
	return
}

func (u *User) GetOne() (user *User, err error) {
	return
}

func (u *User) Update(user *User) (err error) {
	collection := mongodb.Mongodb.Collection("users")
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func (u *User) Delete(id int) (err error) {
	return
}

func (u *User) TranDemo() (err error) {
	return err
}

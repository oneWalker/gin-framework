package mongodb

import (
	"context"
	"gin-practice/pkg/db/mongodb"
	"reflect"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Name     string             `bson:"name" json:"name"`
	Gender   string             `bson:"gender" json:"gender"`
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

//需要注意mongodb的都需要将其转化为相应的流进行解析

func (u *User) GetAll() (users []*User, err error) {
	collection := mongodb.Mongodb.Collection("users")
	// 查询多个
	// 将选项传递给Find()
	findOptions := options.Find()
	findOptions.SetLimit(2)

	// 定义一个切片用来存储查询结果
	var results []*User

	// 把bson.D{{}}作为一个filter来匹配所有文档
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		logrus.Fatal(err)
	}

	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cur.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			logrus.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		logrus.Fatal(err)
	}

	// 完成后关闭游标
	cur.Close(context.TODO())
	logrus.Info("Found multiple documents (array of pointers): %#v\n", results)
	return results, err
}

func (u *User) GetOne(id string) (user *User, err error) {
	// create a value into which the result can be decoded
	var result *User

	//string转为ObjectId
	queryId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": queryId}
	collection := mongodb.Mongodb.Collection("users")
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		logrus.Fatal(err)
	}
	return result, err
}

func (u *User) Update(filter bson.D, user *User) (res *User, err error) {
	collection := mongodb.Mongodb.Collection("users")
	updateResult, err := collection.UpdateOne(context.TODO(), filter, user)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("%v", updateResult.MatchedCount, updateResult.ModifiedCount)
	logrus.Info("%v", user)
	return user, err
}

func (u *User) Delete(id string) (err error) {
	collection := mongodb.Mongodb.Collection("users")

	//多个数据的时候可以使用bson.D,单个数据的时候使用bson.M
	// 删除名字是小黄的那个
	deleteResult1, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		logrus.Fatal(err)
	}
	deleteResult2, err := collection.DeleteMany(context.TODO(), bson.M{"name": "test"})
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("deleted one %v, deleted number %v", deleteResult1.DeletedCount, deleteResult2.DeletedCount)
	return err
}

func (u *User) TranDemo() (err error) {
	ctx := context.Background()
	col := mongodb.Mongodb.Collection("users")
	//定义事务
	session, err := mongodb.Mongodb.Client().StartSession()
	if err != nil {
		logrus.Fatal(err)
		return
	}
	//defer session.EndSession();
	//开启事务
	err = session.StartTransaction()
	if err != nil {
		logrus.Fatal(err)
		return
	}
	//在事务内写一条id为“222”的记录
	_, err = col.InsertOne(ctx, bson.M{"_id": "222", "name": "ddd", "age": 50})
	if err != nil {
		logrus.Fatal(err)
		return
	}

	//写重复id
	_, err = col.InsertOne(ctx, bson.M{"_id": "111", "name": "ddd", "age": 50})
	if err != nil {
		session.AbortTransaction(ctx)
	} else {
		session.CommitTransaction(ctx)
	}
	return err
}

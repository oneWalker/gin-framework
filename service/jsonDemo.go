package service

import (
	"encoding/json"
	"io"
	"reflect"
	"runtime"
	"strings"

	"github.com/goinggo/mapstructure"
	"github.com/sirupsen/logrus"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

//打印当前状态的一种方法
func printMyName() string {
	funcName, file, line, ok := runtime.Caller(1)
	logrus.Info("名字：", funcName, "文件位置：", file, "函数文本位置：", line, ok)
	return runtime.FuncForPC(funcName).Name()
}

//简单方法
func JsonToStruct() {
	logrus.Info(printMyName())
	jsonStr := `{"name":"test","password":"123456"}`
	var user User
	json.Unmarshal([]byte(jsonStr), &user)
	logrus.Info(user)
}

//通过io流的方式进行相关的读取
func JsonToStruct1() {
	logrus.Info(printMyName())
	jsonStr := `{"name":"test","password":"123456"}`
	var user User
	//解析json到结构体，指定长度
	userReader := io.LimitReader(strings.NewReader(jsonStr), 1024)
	errs := json.NewDecoder(userReader).Decode(&user)
	if errs != nil {
		logrus.Info(errs.Error())
	}
	logrus.Info(user)
}

func StructToJson() {
	logrus.Info(printMyName())
	user := User{
		Name:     "test",
		Password: "123456",
	}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info(string(jsonBytes))
}

func MapToStruct() (res User) {
	logrus.Info(printMyName())
	userMap := make(map[string]interface{})
	userMap["Name"] = "test"
	userMap["Password"] = "123456"

	var user User
	err := mapstructure.Decode(userMap, &user)
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info(user)
	return user
}

func StructToMap(obj interface{}) map[string]interface{} {
	logrus.Info(printMyName())
	keys := reflect.TypeOf(obj)
	values := reflect.ValueOf(obj)
	logrus.Info("keys：", keys, " ", "values：", values)
	var data = make(map[string]interface{})
	for i := 0; i < keys.NumField(); i++ {
		data[keys.Field(i).Name] = values.Field(i).Interface()
	}
	logrus.Info(data)
	return data
}

func JsonToMap() {
	logrus.Info(printMyName())
	jsonStr := `{"name":"test","password":"12345"}`
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info(mapResult)
}

func MapToJson() {
	logrus.Info(printMyName())
	userMap := []map[string]interface{}{}
	userMap1 := map[string]interface{}{"name": "test1", "password": "123456"}
	userMap2 := map[string]interface{}{"name": "test2", "password": "123456"}
	userMap = append(userMap, userMap1, userMap2)

	jsonStr, err := json.Marshal(userMap)
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info(string(jsonStr))
}

//参考：https://github.com/emacampolo/gomparator/blob/master/json_util.go#L10
//一个简单的json比较大小，都为json字节流的情况下，比较大小
func JSONIsEqual(j1 []byte, j2 []byte) (bool, error) {
	var i1, i2 interface{}

	if err := json.Unmarshal(j1, &i1); err != nil {
		return false, err
	}

	if err := json.Unmarshal(j2, &i2); err != nil {
		return false, err
	}

	return reflect.DeepEqual(i1, i2), nil
}

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
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

//print current function name
func printMyName() string {
	funcName, file, line, ok := runtime.Caller(1)
	logrus.Info("name：", funcName, "file location：", file, "the current function's file name ：", line, ok)
	return runtime.FuncForPC(funcName).Name()
}

//just use exsited function
func JsonToStruct1(jsonStr string) User {
	var user User
	json.Unmarshal([]byte(jsonStr), &user)
	logrus.Info(user)
	return user
}

//json to struct wih io streaming way
func JsonToStruct2(jsonStr string) User {
	var user User
	//decode json to struct with fixed size
	userReader := io.LimitReader(strings.NewReader(jsonStr), 1024)
	errs := json.NewDecoder(userReader).Decode(&user)
	if errs != nil {
		logrus.Info(errs.Error())
	}
	logrus.Info(user)
	return user
}

func StructToJson(user User) []byte {
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info(string(jsonBytes))
	return jsonBytes
}

func MapToStruct(userMap map[string]interface{}) (res User) {
	var user User
	err := mapstructure.Decode(userMap, &user)
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info(user)
	return user
}

func StructToMap(obj interface{}) map[string]interface{} {
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

func JsonToMap(jsonStr string) map[string]interface{} {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info(mapResult)
	return mapResult
}

// support multi-dimensional map to multi-level json
func MapToJson(userMap map[string]interface{}) []byte {
	//extend map to 2 dimension
	// userMap := []map[string]interface{}{}
	// userMap1 := map[string]interface{}{"name": "test1", "password": "123456"}
	// userMap2 := map[string]interface{}{"name": "test2", "password": "123456"}
	// userMap = append(userMap, userMap1, userMap2)
	objson, err := json.Marshal(userMap)
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info(string(objson))
	return objson
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

package service

//package main

import (
	"encoding/json"
	"encoding/xml"

	"github.com/sirupsen/logrus"
)

type xmlStruct struct {
	XMLName xml.Name `xml:"note"`
	To      string   `xml:"to"`
	From    string   `xml:"from"`
	Heading string   `xml:"heading"`
	Body    string   `xml:"body"`
}

func XmlToStruct() (xmlRes xmlStruct) {
	xmlDoc := `<?xml version="1.0" encoding="UTF-8"?>
                <note>
                  <to>Tove</to>
                  <from>Jani</from>
                  <heading>Reminder</heading>
                  <body>Don't forget me this weekend!</body>
                </note>`
	x := xmlStruct{}
	err := xml.Unmarshal([]byte(xmlDoc), &x)
	if err != nil {
		logrus.Fatalf(err.Error())
	} else {
		logrus.Info(x.Body)
	}
	return x
	// res, _ := json.Marshal(x)
	// logrus.Info(reflect.TypeOf(res))
	// logrus.Info(reflect.TypeOf(x))
	// var mapResult map[string]interface{}
	// json.Unmarshal([]byte(res), &mapResult)
	// logrus.Info(mapResult)

}

func StructToXml() {
	//struct to xml
	sd := xmlStruct{To: "Pika", From: "Brian", Heading: "No Subject", Body: "I love you"}
	//struct添加新对象,使用append

	sdXml, _ := xml.Marshal(&sd)
	//通过MarshalIndent，让xml数据输出好看点
	//sdXml, _ := xml.MarshalIndent(&sd, "", "  ")
	//得到的sdXml都是字节流
	logrus.Info(sdXml)
	logrus.Info(string(sdXml))
}

//XmlToJson Directly:https://github.com/basgys/goxml2json
// the demo provide the way by the Struct
func XmlToJson() (res []uint8) {
	//xml to struct
	x := XmlToStruct()
	//struct to json
	res, _ = json.Marshal(x)
	logrus.Info(string(res))
	return res
}

func JsonToXml() {
	x := XmlToJson()
	var f xmlStruct
	//json str to struct
	json.Unmarshal(x, &f)
	logrus.Info("json")
	logrus.Info(f)
	//stuct to xml
	sdXml, _ := xml.Marshal(&f)
	logrus.Info(sdXml)
	logrus.Info(string(sdXml))
}

// main函数用于测试，需要运行使用时需要引入package main
// func main() {
// 	XmlToStruct()
// 	StructToXml()
// 	XmlToJson()
// 	JsonToXml()
// }

package service

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

func XmlToStruct(xmlDoc string) xmlStruct {
	x := &xmlStruct{}
	err := xml.Unmarshal([]byte(xmlDoc), x)
	if err != nil {
		logrus.Fatalf(err.Error())
	} else {
		logrus.Info(x.XMLName.Local)
	}
	return *x
	// res, _ := json.Marshal(x)
	// logrus.Info(reflect.TypeOf(res))
	// logrus.Info(reflect.TypeOf(x))
	// var mapResult map[string]interface{}
	// json.Unmarshal([]byte(res), &mapResult)
	// logrus.Info(mapResult)

}

func StructToXml(sd xmlStruct) string {
	//struct to xml
	//struct添加新对象,使用append

	sdXml, _ := xml.Marshal(&sd)
	//通过MarshalIndent，让xml数据输出好看点
	//sdXml, _ := xml.MarshalIndent(&sd, "", "  ")
	//得到的sdXml都是字节流
	logrus.Info(sdXml)
	logrus.Info(string(sdXml))
	return string(sdXml)
}

//XmlToJson Directly:https://github.com/basgys/goxml2json
//the demo provide the way by the Struct
//byte is equal to uint8
//string is equal to []byte
func XmlToJson(xml string) []byte {

	//方法1: first xml to struct, then struct to json
	x := XmlToStruct(xml)

	res, _ := json.Marshal(x)

	//方法2: Convert xml to json directly
	// res, _ = goxml2json.XmlToJson([]byte(xml))

	return res
}

func JsonToXml(jsonStr []byte) []byte {
	var f xmlStruct
	//json str to struct
	json.Unmarshal(jsonStr, &f)
	//stuct to xml
	sdXml, _ := xml.Marshal(&f)

	logrus.Info(sdXml)
	logrus.Info(string(sdXml))
	return sdXml
}

func MapToXml(m map[string]string) (res []byte) {
	//convert required the sturct defined in stringmap
	buf, err := xml.Marshal(Map(m))
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(string(buf))
	return buf
}

//xml convert only can support map[string]string cannot support map[string]interface{}

func XmlToMap(buf []byte) map[string]string {
	var rm map[string]string
	err := xml.Unmarshal(buf, (*Map)(&rm))
	if err != nil {
		logrus.Error(err)
	}
	return rm
}

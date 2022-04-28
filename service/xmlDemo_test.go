package service

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"testing"
)

const xmlstring = `<?xml version="1.0" encoding="UTF-8"?><note><To>Tove</To><From>Jani</From><Heading>Reminder</Heading><Body>Dont_forget_me_this_weekend!</Body></note>`
const Header = `<?xml version="1.0" encoding="UTF-8"?>`

var sd = xmlStruct{To: `Tove`, From: `Jani`, Heading: `Reminder`, Body: `Dont_forget_me_this_weekend!`}

var jsonStr = `{"To":"Tove","From":"Jani","Heading":"Reminder","Body":"Dont_forget_me_this_weekend!"}`

var mapStr = map[string]string{
	"To":      "Tove",
	"From":    "Jani",
	"Heading": "Reminder",
	"Body":    "Dont_forget_me_this_weekend!",
}

func TestXmlToStruct(t *testing.T) {
	type args struct {
		xmlDoc string
	}
	tests := []struct {
		name       string
		args       args
		wantXmlRes xmlStruct
	}{
		{
			name:       "TestXmlToStruct",
			args:       args{xmlstring},
			wantXmlRes: sd,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//generate structs时无法将其转化为相应的或者删除xmltoStruct中的xml的值,转换为json的时候可以转换到相应的note值
			tt.wantXmlRes.XMLName = xml.Name{Local: "note"}

			if gotXmlRes := XmlToStruct(tt.args.xmlDoc); !reflect.DeepEqual(gotXmlRes, tt.wantXmlRes) {
				t.Logf(gotXmlRes.XMLName.Local)
				t.Log(tt.wantXmlRes.XMLName.Local)
				t.Errorf("XmlToStruct() = %v, want %v", gotXmlRes, tt.wantXmlRes)
			}
		})
	}
}

func TestStructToXml(t *testing.T) {
	type args struct {
		sd xmlStruct
	}
	tests := []struct {
		name       string
		args       args
		wantXmlDoc string
	}{
		{
			name:       "TestStructToXml",
			args:       args{sd},
			wantXmlDoc: xmlstring,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//'号在不同字符下会产生不同，因此报错
			if gotXmlDoc := Header + StructToXml(tt.args.sd); gotXmlDoc != tt.wantXmlDoc {
				t.Errorf("StructToXml() = %v, want %v", gotXmlDoc, tt.wantXmlDoc)
			}
		})
	}
}

func TestXmlToJson(t *testing.T) {
	type args struct {
		xml string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []byte
	}{
		{
			name:    "TestXmlToJson",
			args:    args{xmlstring},
			wantRes: []byte(jsonStr),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//原生的json和struct都只有在变为map时才可以删除相关的key，相互比较也需要转化为map进行比较
			gotRes := XmlToJson(tt.args.xml)

			var m map[string]interface{}
			if err := json.Unmarshal([]byte(gotRes), &m); err != nil {
				panic(err)
			}
			if _, exists := m["XMLName"]; exists {
				delete(m, "XMLName")
				gotResAfter, err := json.Marshal(m)
				if err != nil {
					panic(err)
				}
				gotRes = gotResAfter
			}

			var j1, j2 interface{}
			json.Unmarshal(gotRes, &j1)
			json.Unmarshal(tt.wantRes, &j2)

			if !reflect.DeepEqual(j1, j2) {
				t.Errorf("XmlToJson() = %v, want %v", j1, j2)
			}
		})
	}
}

func TestJsonToXml(t *testing.T) {
	type args struct {
		jsonStr []byte
	}
	tests := []struct {
		name    string
		args    args
		wantRes []byte
	}{
		{
			name:    "TestJsonToXml",
			args:    args{[]byte(jsonStr)},
			wantRes: []byte(xmlstring),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//json在marshal的过程会将key小写化

			if gotRes := Header + string(JsonToXml(tt.args.jsonStr)); !reflect.DeepEqual(gotRes, string(tt.wantRes)) {
				t.Errorf("JsonToXml() = %v, want %v", string(gotRes), string(tt.wantRes))
			}
		})
	}
}

func TestMapToXml(t *testing.T) {
	type args struct {
		m map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantRes []byte
	}{
		{
			name:    "TestMapToXml",
			args:    args{mapStr},
			wantRes: []byte(xmlstring),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := MapToXml(tt.args.m); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("MapToXml() = %v, want %v", string(gotRes), string(tt.wantRes))
			}
		})
	}
}

func TestXmlToMap(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		args    args
		wantRes map[string]string
	}{
		{
			name:    "TestXmlToMap",
			args:    args{[]byte(xmlstring)},
			wantRes: mapStr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := XmlToMap(tt.args.buf); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("MapToXml() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

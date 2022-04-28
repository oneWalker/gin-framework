package service

import (
	"reflect"
	"testing"
)

const userStr = `{"Name":"test","Password":"123456"}`

var userRes = User{"test", "123456"}

var userMap = map[string]interface{}{
	"Name":     "test",
	"Password": "123456",
}

func TestJsonToStruct1(t *testing.T) {
	type args struct {
		jsonStr string
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		{
			name: "TestJsonToStruct1",
			args: args{userStr},
			want: userRes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JsonToStruct1(tt.args.jsonStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToStruct1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printMyName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test_printMyName",
			want: "gin-practice/service.Test_printMyName.func1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printMyName(); got != tt.want {
				t.Errorf("printMyName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonToStruct2(t *testing.T) {
	type args struct {
		jsonStr string
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		{
			name: "TestJsonToStruct1",
			args: args{userStr},
			want: userRes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JsonToStruct2(tt.args.jsonStr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToStruct2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStructToJson(t *testing.T) {
	type args struct {
		user User
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "TestStructToJson",
			args: args{userRes},
			want: []byte(userStr),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructToJson(tt.args.user); !reflect.DeepEqual(string(got), string(tt.want)) {
				t.Errorf("StructToJson() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestMapToStruct(t *testing.T) {
	type args struct {
		userMap map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantRes User
	}{
		{
			name:    "TestMapToStruct",
			args:    args{userMap},
			wantRes: userRes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := MapToStruct(tt.args.userMap); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("MapToStruct() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestStructToMap(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "TestStructToMap",
			args: args{userRes},
			want: userMap,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StructToMap(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

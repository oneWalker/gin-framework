package controller

import (
	"fmt"
	"gin-practice/model/request"
	"gin-practice/model/response"
	"github.com/gin-gonic/gin"
	"net/http"

	"io/ioutil"
)

//query
func GetTest(ctx *gin.Context) {

	var T request.Login
	//GET All params value
	Params := ctx.Request.URL.Query()
	T.Password = Params.Get("password")
	T.Username = Params.Get("username")

	// //GET pointed param value
	// T.Username = ctx.Query("username")

	// //GET pointed param value and set a deault value when it doesnt exist
	// T.Username = ctx.DefaultQuery("username", "default")

	response.SUCCESS(T, ctx)
}

//testform struct/raw
func PostTestForm(ctx *gin.Context) {

	//通过结构体进行获取
	var T request.Login
	if err := ctx.Bind(&T); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//单个获取
	T.Username = ctx.PostForm("username")
	T.Password = ctx.PostForm("password")
	// 多选框
	T.Hobbys = ctx.PostFormArray("hobby")

	response.SUCCESS(T, ctx)
}

//test json
func PostTestJson(ctx *gin.Context) {

	//使用于application/json
	var T request.Login
	if err := ctx.ShouldBindJSON(&T); err != nil {
		fmt.Println("error", err)
		//return
	}

	//只返回相应的值但并不做相关的处理
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	fmt.Printf("ctx.Request.body: %v", string(data))

	//fmt.Println("params", T)
	response.SUCCESS(T, ctx)
}

func TestUri(ctx *gin.Context) {

	var T request.Login
	//结构体匹配
	if err := ctx.ShouldBindUri(&T); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//单独获取
	name := ctx.Param("name")
	password := ctx.Param("password")
	fmt.Println(name, password)
	// 判断用户名密码是否正确
	if T.Username != "root" || T.Password != "admin" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "200"})
}

//test file
func PostTestFile(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	ctx.SaveUploadedFile(file, file.Filename)
	ctx.JSON(http.StatusOK, gin.H{"status": "200", "filename": file.Filename})
}

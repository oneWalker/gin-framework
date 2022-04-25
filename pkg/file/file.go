package file

import (
	"fmt"
	"gin-practice/pkg/core"
	"io"
	"os"
	"strings"

	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// UploadFileToCos 上传文件到Cos
func UploadFileToCos(filePath, fileName string) (string, error) {
	cosURL := viper.GetString("tencentCOS.uploadFile")
	uploadReq := req.New()
	uploadReq.SetFlags(req.LreqHead) // output request head (request line and request header)
	uploadHeader := req.Header{
		"Content-Type": "multipart/form-data",
	}
	uploadParam := req.Param{
		"file_path":       filePath,
		"allow_duplicate": "1",
	}
	file, err := os.Open(fileName)
	defer os.Remove(fileName)
	uploadFile := req.FileUpload{
		FileName:  fileName,
		FieldName: "file",
		File:      file,
	}
	resp, err := uploadReq.Post(cosURL, uploadHeader, uploadParam, uploadFile)
	if err != nil {
		logrus.Errorf("resp.Post() error: %s", err)
		return "", err
	}
	// 返回的内容
	type responseObject struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			FileName        string `json:"file_name"`
			FileUrl         string `json:"file_url"`
			FileNameWithCDN string `json:"file_name_with_cdn"`
		} `json:"data"`
	}
	cosResponse := new(responseObject)
	if err := resp.ToJSON(cosResponse); err != nil {
		logrus.Errorf("resp.ToJSON() error: %s", err)
		return "", err
	}
	return cosResponse.Data.FileNameWithCDN, nil
}

// UploadByteToCos 上传流到Cos文件
func UploadByteToCos(filePath, fileName string, file io.ReadCloser, ctx *core.Context) (string, error) {
	cosURL := viper.GetString("tencentCOS.uploadFile")
	uploadReq := req.New()
	uploadReq.SetFlags(req.LstdFlags) // output request head (request line and request header)
	uploadHeader := req.Header{
		"Content-Type":  "multipart/form-data",
		"Authorization": ctx.Request.Header.Get("Authorization"),
	}
	uploadParam := req.Param{
		"file_path": filePath,
		//"allow_duplicate": "1",
	}
	uploadFile := req.FileUpload{
		FileName:  fileName,
		FieldName: "file",
		File:      file,
	}
	resp, err := uploadReq.Post(cosURL, uploadHeader, uploadParam, uploadFile)
	if err != nil {
		logrus.Errorf("resp.Post() error: %s", err)
		return "", err
	}
	// 返回的内容
	type responseObject struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			FileName        string `json:"file_name"`
			FileUrl         string `json:"file_url"`
			FileNameWithCDN string `json:"file_name_with_cdn"`
		} `json:"data"`
	}
	cosResponse := new(responseObject)
	if err := resp.ToJSON(cosResponse); err != nil {
		logrus.Errorf("resp.ToJSON() error: %s", err)
		return "", err
	}
	return cosResponse.Data.FileNameWithCDN, nil
}

func FileNameEscape(uniquenessName string) string {
	charList := GetCharList()
	// 获取文件名
	filePathArray := strings.Split(uniquenessName, "/")
	filename := filePathArray[len(filePathArray)-1]
	// 编码文件名
	filenameArray := strings.Split(filename, "")
	for k, v := range filenameArray {
		for _, char := range charList {
			if v == char {
				logrus.Infof("the %s char is :%s\n", uniquenessName, char)
				filenameArray[k] = "-"
				break
			}
		}
	}
	path := strings.Join(filePathArray[:len(filePathArray)-1], "/")
	uniquenessName = fmt.Sprintf("%s%s", path, strings.Join(filenameArray, ""))
	return uniquenessName
}

func GetCharList() []string {
	var results []string
	// 英文字符
	results = append(results, ",", ":", ";", "=", "&", "$", "@", "+", "?", " ")
	// 中文字符
	results = append(results, "，", "：", "；", "？")
	// ASCII 字符范围：00-1F 十六进制（0-31 十进制）以及7F（127 十进制）
	for i := 0; i <= 31; i++ {
		runei := rune(i)
		results = append(results, string(runei))
	}
	results = append(results, string(rune(127)))
	// 尽量避免使用的特殊字符
	// 英文
	results = append(results, "`", "^", "\"", "\\", "{", "}", "[", "]", "~", "%", "#", "|")
	// 中文
	results = append(results, "【", "】", "《", "》")
	// 尽量避免使用的ASCII 128-255 十进制
	for i := 128; i <= 255; i++ {
		runei := rune(i)
		results = append(results, string(runei))
	}
	return results
}

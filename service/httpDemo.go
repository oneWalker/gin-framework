package service

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//一个通用的get和post的http请求
//返回相应的json数据
//其他方式可使用：http.Get，http.PostForm，http.Post
func HttpCommon(host string, data map[string]string, method string, headers map[string]string) (result map[string]string, err error) {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	var urlpath string
	var inputIO io.Reader
	//默认的headers为application/json
	hasContent := false
	client := &http.Client{}
	switch method {
	case "get":
		params := url.Values{}
		urlObj, _ := url.Parse(host)
		for k, v := range data {
			params.Set(k, v)
		}
		urlObj.RawQuery = params.Encode()
		urlpath = urlObj.String()
	case "post":
		urlpath = host
		//return []byte
		//方式2，使用get方式进行编码（不对host进行Parse），然后传入strings.NewReader(params)
		bytesData, _ := json.Marshal(data)
		inputIO = bytes.NewReader(bytesData)
	}
	req, _ := http.NewRequest(method, urlpath, inputIO)

	for k, v := range headers {
		if strings.ToLower(k) == "content-type" {
			hasContent = true
		}
		req.Header.Add(k, v)
	}

	//header中没有设置默认content的时候，设置一个默认值
	if !hasContent {
		req.Header.Add("Content", "application/json")
	}

	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	res := make(map[string]string)

	//json([]byte) to map
	//其中的res也可以是struct类型的结构
	err = json.Unmarshal(body, &res)

	return res, err
}

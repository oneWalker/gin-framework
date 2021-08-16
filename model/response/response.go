package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
	c.Abort()
}

func SUCCESS(data interface{}, c *gin.Context) {
	Result(0, data, "success", c)
}

func ERROR(message string, c *gin.Context) {
	Result(-1, map[string]interface{}{}, message, c)
}

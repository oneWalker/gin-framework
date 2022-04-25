package core

import (
	"gin-practice/pkg/errno"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Context struct.
type Context struct {
	*gin.Context
}

// Success method.
func (c *Context) Success(data interface{}) {

	ret := gin.H{
		"code": errno.OK.Code,
		"msg":  errno.OK.Message,
		"data": data,
	}

	c.JSON(http.StatusOK, ret)
}

// Fail method.
func (c *Context) Fail(err error, httpStatusCodes ...int) {
	logrus.Errorf("%+v", err)
	ret := gin.H{}
	if err != nil {
		if errCode, ok := err.(*errno.Errno); ok {
			ret["code"] = errCode.Code
			ret["msg"] = errCode.Message
		} else {
			ret["code"] = errno.FAILED.Code
			ret["msg"] = err.Error()
		}
	} else {
		ret["code"] = errno.FAILED.Code
		ret["msg"] = errno.FAILED.Message
	}

	code := http.StatusBadRequest
	if len(httpStatusCodes) != 0 {
		code = httpStatusCodes[0]
	}

	c.AbortWithStatusJSON(code, ret)
}

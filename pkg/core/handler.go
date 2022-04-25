package core

import (
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(*Context)

// Handle transform HandlerFuc to gin.HandlerFunc
func Handle(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{
			Context: c,
		}
		handler(ctx)
	}
}

func Ping(c *Context) {
	c.Success("ping")
}

package middleware

import (
	"gin-practice/pkg/core"
	"gin-practice/pkg/errno"
	"gin-practice/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth .
func JWTAuth() gin.HandlerFunc {

	return func(context *gin.Context) {
		ctx := &core.Context{Context: context}
		auth := ctx.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			ctx.Fail(errno.TokenMissing, http.StatusUnauthorized)
			return
		}
		tokenEncode := strings.Fields(auth)[0]

		// 校验token
		token, err := jwt.NewJWT().ParseToken(tokenEncode)
		if err != nil {
			ctx.Fail(errno.TokenExpired, http.StatusUnauthorized)
			return
		}

		ctx.Set("token", token)
		ctx.Next()
	}
}

package http

import (
	"easy_im/internal/api/pkg"
	"easy_im/pkg/jwt"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Request.Header.Get(pkg.HeaderToken)
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusOK, pkg.UserError(errors.New("without token")))
			return
		}
		// 校验token，只要出错直接拒绝请求
		c, err := jwt.DecodeToken(auth)
		if err != nil {
			context.Abort()
			context.JSON(http.StatusOK, pkg.UserError(err))
			return
		}
		// 续签
		token, err := jwt.RenewToken(c)
		if err == nil {
			context.Request.Header.Set(pkg.HeaderToken, token)
		}
		context.Set(pkg.CTXUserId, c.UserId)
		context.Set(pkg.CTXUserName, c.UserName)
		context.Next()
	}
}

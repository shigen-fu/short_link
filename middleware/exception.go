package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
	"short-link/base/errno"
)

func ExceptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				// 修改错误信息解析
				code, err := convertErr(e)
				c.JSON(code, err)
				c.Abort()
			}
		}()
		c.Next()
	}
}

func convertErr(err interface{}) (code int, e errno.Error) {
	switch v := err.(type) {
	case errno.Error:
		return http.StatusOK, v
	case error:
		debug.PrintStack()
		log.Printf("panic: %v\n", v.Error())
		return http.StatusInternalServerError, errno.ErrServer.WithData("服务器发生错误")
	default:
		return http.StatusInternalServerError, errno.ErrUnknownServer
	}
}

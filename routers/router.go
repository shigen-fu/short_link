package routers

import (
	"github.com/gin-gonic/gin"
	handler "short-link/handler/short_link"
	"short-link/middleware"
	link "short-link/service/short_link"
)

var (
	shortLinkHandler *handler.ShortLinkHandler
)

func Init(r *gin.Engine) {
	initService()
	initHandler()
	initShortLinkRouter(r)
}

func initService() {
	link.InitService()
}

func initHandler() {
	shortLinkHandler = handler.InitShortLinkHandler()
}

func initShortLinkRouter(r *gin.Engine) {
	r.Use(middleware.Cors())
	r.Use(middleware.ExceptionMiddleware())
	groupV1 := r.Group("/v1")
	{
		groupV1.GET("/add", shortLinkHandler.Add)
		groupV1.GET("/get", shortLinkHandler.Get)
	}
}

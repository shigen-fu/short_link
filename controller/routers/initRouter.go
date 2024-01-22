package routers

import (
	"github.com/gin-gonic/gin"
	"short-link/controller/handlers"
)

func InitRouter(r *gin.Engine) {
	groupV1 := r.Group("/v1")
	{
		groupV1.GET("/add", handlers.Create)
		groupV1.GET("/get", handlers.Get)
	}
}

package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"short-link/base/config"
	"short-link/base/db"
	"short-link/middleware"
	"short-link/routers"
)

var path = flag.String("p", "./conf/db.ini", "")
var addr = flag.String("s", "127.0.0.1:8099", "server addr")

func main() {
	flag.Parse()
	config.Init(*path)
	db.Init()

	r := gin.Default()
	r.Use(middleware.ExceptionMiddleware())
	routers.Init(r)
	r.Run(*addr)
}

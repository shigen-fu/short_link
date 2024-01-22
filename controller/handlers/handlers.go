package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"short-link/base/db"
	"short-link/base/errno"
	"short-link/base/utils"
	"short-link/model"
	"time"
)

func Create(c *gin.Context) {
	originLink := c.Query("link")
	if len(originLink) == 0 {
		c.JSON(http.StatusOK, errno.ErrLink)
	}
	code := utils.GenerateCode(originLink)
	link := model.ShortLink{
		Id:         0,
		OriginURL:  originLink,
		ShortCode:  code,
		CreateTime: time.Now().Unix(),
		Creator:    0,
	}
	// &是希望回写id
	db.GetMysqlDb().Create(&link)
	// save to redis
	db.GetRedisDb().Set(c, code, originLink, time.Second*60*10)
	c.JSON(http.StatusOK, errno.OK.WithData(link))
}

func Get(c *gin.Context) {
	code := c.Query("code")
	var shortLink model.ShortLink
	if len(code) == 0 {
		c.JSON(http.StatusOK, errno.ErrLink)
	}
	// from redis
	originLink, err := db.GetRedisDb().Get(c, code).Result()
	if err == redis.Nil {
		// from mysql
		if err = db.GetMysqlDb().Where("short_code = ?", code).Find(&shortLink).Error; err != nil {
			fmt.Printf("mysql err: %#v\n", err)
			c.JSON(http.StatusOK, errno.ErrServer)
			return
		}
		db.GetRedisDb().Set(c, code, shortLink.OriginURL, time.Second*60*10)
		c.JSON(http.StatusOK, errno.OK.WithData(shortLink.OriginURL))
		return
	} else if err != nil {
		fmt.Printf("redis err: %#v\n", err)
		c.JSON(http.StatusOK, errno.ErrServer)
		return
	} else {
		c.JSON(http.StatusOK, errno.OK.WithData(originLink))
	}
}

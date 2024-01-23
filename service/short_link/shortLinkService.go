package short_link

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"short-link/base/db"
	"short-link/base/errno"
	"short-link/base/utils"
	"short-link/model"
	"short-link/service"
	"sync"
	"time"
)

var (
	once sync.Once
)

func InitService() {
	once.Do(func() {
		service.ShortLinkService = &userService{}
	})
}

type userService struct{}

func (u userService) Add(link string) (shortLink model.ShortLink, err error) {
	if len(link) == 0 {
		panic(interface{}(errno.ErrLink))
	}

	code := utils.GenerateCode(link)
	shortLinkModel := model.ShortLink{
		Id:         0,
		OriginURL:  link,
		ShortCode:  code,
		CreateTime: time.Now().Unix(),
		Creator:    0,
	}
	// &是希望回写id
	db.GetMysqlDb().Create(&link)
	// save to redis
	db.GetRedisDb().Set(context.Background(), code, link, time.Second*60*10)
	return shortLinkModel, nil
}

func (u userService) Get(code string) (originUrl string, err error) {
	var shortLink model.ShortLink
	c := context.Background()
	if len(code) == 0 {
		panic(interface{}(errno.ErrParam))
	}
	// from redis
	originLink, err := db.GetRedisDb().Get(c, code).Result()
	if err == redis.Nil {
		// from mysql
		if count := db.GetMysqlDb().Where("short_code = ?", code).Find(&shortLink).RowsAffected; count == 0 {
			fmt.Printf("code: %#v has no refer link\n", code)
			panic(interface{}(errno.ErrLinkNotExist))
		}
		db.GetRedisDb().Set(c, code, shortLink.OriginURL, time.Second*60*10)
		return originLink, nil
	} else if err != nil {
		fmt.Printf("redis err: %#v\n", err)
		panic(interface{}(errno.ErrServer))
	} else {
		return originLink, nil
	}
}

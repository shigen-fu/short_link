package service

import "short-link/model"

var (
	ShortLinkService IShortLinkService
)

func GetShortLinkService() IShortLinkService {
	return ShortLinkService
}

type IShortLinkService interface {
	Add(link string) (shortLink model.ShortLink, err error)
	Get(code string) (originUrl string, err error)
}

package short_link

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"short-link/base/errno"
	"short-link/service"
)

type ShortLinkHandler struct {
	shortLinkService service.IShortLinkService
}

func InitShortLinkHandler() *ShortLinkHandler {
	return &ShortLinkHandler{shortLinkService: service.GetShortLinkService()}
}

func (s ShortLinkHandler) Add(c *gin.Context) {
	link := c.Query("link")
	shortLinkModel, _ := s.shortLinkService.Add(link)
	c.JSON(http.StatusOK, errno.OK.WithData(shortLinkModel))
}

func (s ShortLinkHandler) Get(c *gin.Context) {
	code := c.Query("code")
	originUrl, _ := s.shortLinkService.Get(code)
	c.JSON(http.StatusOK, errno.OK.WithData(originUrl))
}

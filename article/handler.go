package article

import (
	"github.com/gin-gonic/gin"
)

type HandlerService struct {
	service *Service
}

type HandlerInterface interface {
	postArticle(ctx *gin.Context)
	getArticle(ctx *gin.Context)
}

func makeHandler(router *gin.RouterGroup, service *Service) {
	handler := &HandlerService{
		service: service,
	}
	router.POST("/insertArticle", handler.postArticle)
	router.GET("/getArticle/:id", handler.getArticle)
}

func (handler *HandlerService) postArticle(ctx *gin.Context) {
	handler.service.postArticleInService(ctx)
}

func (handler *HandlerService) getArticle(ctx *gin.Context) {
	handler.service.getArticleInService(ctx)
}

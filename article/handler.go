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

func (service *Service) newHandler() *HandlerService {
	return &HandlerService{
		service: service,
	}

}
func (handler *HandlerService) makeHandler(router *gin.RouterGroup) {

	router.POST("/insertArticle", handler.postArticle)
	router.GET("/getArticle/:id", handler.getArticle)
}

func (handler *HandlerService) postArticle(ctx *gin.Context) {
	handler.service.postArticleInService(ctx)
}

func (handler *HandlerService) getArticle(ctx *gin.Context) {
	handler.service.getArticleInService(ctx)
}

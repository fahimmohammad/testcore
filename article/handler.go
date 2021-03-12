package article

import (
	"github.com/gin-gonic/gin"
)

type HandlerService struct {
	service *Service
}

// type HandlerInterface interface {
// 	PostArticle(ctx *gin.Context)
// 	GetArticle(ctx *gin.Context)
// }

func (service *Service) GetHandler() *HandlerService {
	return &HandlerService{
		service: service,
	}

}
func (handler *HandlerService) MakeHandler(router *gin.RouterGroup) {

	router.POST("/insertArticle", handler.PostArticle)
	router.GET("/getArticle/:id", handler.GetArticle)
}

func (handler *HandlerService) PostArticle(ctx *gin.Context) {
	handler.service.PostArticleInService(ctx)
}

func (handler *HandlerService) GetArticle(ctx *gin.Context) {
	handler.service.GetArticleInService(ctx)
}

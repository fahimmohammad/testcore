package article

import (
	"net/http"
	"strconv"
	"time"

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
	router.POST("/Article", handler.postArticle)
	router.GET("/Article/:id", handler.getArticle)
	router.PUT("/Article/:id", handler.updateArticle)
	router.DELETE("/Article/:id", handler.deleteArticle)
}

func (handler *HandlerService) postArticle(ctx *gin.Context) {
	var article Article
	err := ctx.ShouldBindJSON(&article)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
	}
	article.CreatedAt = time.Now().String()
	article, err = handler.service.PostArticleInService(article)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, article)
}
func (handler *HandlerService) getArticle(ctx *gin.Context) {
	var article Article
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Not an integer")
	}
	article, err = handler.service.GetArticleInService(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": err})
		return
	} else {
		ctx.JSON(http.StatusOK, article)
	}
}

func (handler *HandlerService) updateArticle(ctx *gin.Context) {
	var article Article
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Not an integer")
	}

	err = ctx.ShouldBindJSON(&article)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
	}
	article, err = handler.service.UpdateArticleInService(id, article)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": err})
		return
	} else {
		ctx.JSON(http.StatusOK, article)
	}

}

func (handler *HandlerService) deleteArticle(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Not an integer")
	}
	err = handler.service.DeleteArticleInService(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted"})
	return
}

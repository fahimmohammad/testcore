package article

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Service struct {
	db *DBStruct
}

type ServiceInterface interface {
	PostArticleInService()
	GetArticleInService()
}

func NewService(dbs *DBStruct) *Service {
	return &Service{
		db: dbs,
	}
}

func (service *Service) PostArticleInService(ctx *gin.Context) {
	var article Article
	err := ctx.ShouldBindJSON(&article)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, article)
	}
	article.CreatedAt = time.Now().String()
	article, err = service.db.InsertArticle(article)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, article)

}

func (service *Service) GetArticleInService(ctx *gin.Context) {
	var article Article
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Not an integer")
	}
	article, err = service.db.FindArticle(id, article)
	if err != nil {
		ctx.JSON(http.StatusNoContent, "Not found")
	}
	ctx.JSON(http.StatusOK, article)

}

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
	postArticleInService()
	getArticleInService()
	newService()
}

func newService(dbs *DBStruct) *Service {
	return &Service{
		db: dbs,
	}
}

func (service *Service) postArticleInService(ctx *gin.Context) {
	var article Article
	err := ctx.ShouldBindJSON(&article)
	if err != nil {

		ctx.JSON(http.StatusOK, err)
	}
	article.CreatedAt = time.Now().String()
	article, err = service.db.insertArticle(article)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, article)

}

func (service *Service) getArticleInService(ctx *gin.Context) {
	var article Article
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Not an integer")
	}
	article, err = service.db.findArticle(id, article)
	//fmt.Println(err)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": err})
		return
	} else {
		ctx.JSON(http.StatusOK, article)
	}

}

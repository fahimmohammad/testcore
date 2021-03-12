package article

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func StartArticle(session *mgo.Session, router *gin.RouterGroup) {
	db := NewDB(session)
	service := NewService(db) // ne
	handler := service.GetHandler()
	handler.MakeHandler(router)

	/*router.POST("/insertArticle", db.PostArticle)
	router.GET("/getArticle/:id", db.GetArticle)*/

}

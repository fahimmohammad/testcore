package article

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func StartArticle(session *mgo.Session, router *gin.RouterGroup) {
	db := newDB(session)
	service := newService(db)
	handler := service.newHandler()
	handler.makeHandler(router)
}

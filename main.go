package main

import (
	"fmt"

	"github.com/fahimsgit/testCore/article"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

type DBSession struct {
	Session   *mgo.Session
	DbName    string
	TableName string
}

type Article struct {
	ID          int32  `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	IsPublished bool   `json:"ispublished" bson:"ispublished"`
}

func main() {
	var err error

	//router := gin.Default()
	router := gin.New()
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("Could not connect")
	}
	api := router.Group("/api/v1/")
	article.StartArticle(session, api)
	/*db := DBSession{
		Session:   session,
		DbName:    "testDb",
		TableName: "article",
	}
	router.POST("/insertArticle", db.postArticle)
	router.GET("/getArticle/:id", db.getArticle)*/

	router.Run(":8080")

}

/*func (d *DBSession) insertArticle(article Article) (Article, error) {

	table := d.Session.DB(d.DbName).C(d.TableName)
	err := table.Insert(article)
	if err != nil {
		return Article{}, err
	} else {
		return article, nil
	}
}

func (d *DBSession) postArticle(ctx *gin.Context) {
	var article Article
	err := ctx.ShouldBindJSON(&article)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, article)
	}
	insertedArticle, err := d.insertArticle(article)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, insertedArticle)

}

func (d *DBSession) findArticle(id int, article Article) (Article, error) {

	err := d.Session.DB(d.DbName).C(d.TableName).Find(bson.M{"id": id}).One(&article)
	if err != nil {

		return Article{}, err
	}
	return article, nil
}

func (d *DBSession) getArticle(ctx *gin.Context) {
	var article Article
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Not an integer")
	}
	article, err = d.findArticle(id, article)
	if err != nil {
		ctx.JSON(http.StatusNoContent, "Not found")
	}
	ctx.JSON(http.StatusOK, article)

}
*/

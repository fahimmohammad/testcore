package main

import (
	"fmt"
	"time"

	"github.com/fahimsgit/testCore/article"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

func main() {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("Could not connect")
	}
	api := router.Group("/api/v1/")
	startService(session, api)
	router.Run(":8080")
}

func startService(session *mgo.Session, api *gin.RouterGroup) {
	article.StartArticle(session, api)
}

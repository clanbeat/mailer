package main

import (
	"fmt"
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/clanbeat/mailer/src/sender"
	"log"
	"os"
)

func initRoutes(templatePath string) {
	log.Println("init router")
	tmplts := createRender(templatePath)
	router := gin.New()
	router.HTMLRender = tmplts
	router.Use(gin.Logger())
	router.Use(RecoveryMiddleware(os.Stdout, errorTracker))

	router.GET("/emails/:name", getEmail)
	router.OPTIONS("/emails/:name")

	router.GET("/test/:name/:email", sendTest)
	router.OPTIONS("/test/:name/:email")

	router.GET("/status", getStatus)
	router.OPTIONS("/status")
	v1 := router.Group("/v1")
	{
		v1.POST("/send", postEmail)
		v1.OPTIONS("/send")
	}
	//start server
	log.Println(fmt.Sprintf("=> Starting in %s on http://0.0.0.0:%s", env.Env, os.Getenv("PORT")))
	router.Run(":" + os.Getenv("PORT"))

}

func createRender(templatePath string) sender.TemplateCache {
	c, err := sender.BuildTemplateCache(templatePath)
	if err != nil {
		fatalError(err)
	}
	return c
}

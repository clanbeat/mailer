package main

import (
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"log"
	"os"
)

func initRoutes() {
	log.Println("init router")
	router = gin.New()
	// Global middlewares
	router.Use(gin.Logger())
	router.Use(RecoveryMiddleware(os.Stdout, errorTracker))

	router.GET("/status", getStatus)
	router.OPTIONS("/status")
}

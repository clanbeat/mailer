package main

import (
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func getStatus(c *gin.Context) {
	// TODO check if storage, message queue or other dependent system is up

	c.JSON(http.StatusOK, gin.H{"status": "OK", "timestamp": time.Now().Unix()})
}

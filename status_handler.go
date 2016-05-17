package main

import (
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK", "timestamp": time.Now().Unix()})
}

package main

import (
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/clanbeat/mailer/src/sender"
	"log"
	"net/http"
)

func postEmail(c *gin.Context) {
	m := &sender.Mailable{}

	if err := c.BindJSON(m); err != nil {
		c.JSON(badRequest(err.Error()))
		return
	}
	log.Println("email", m)
	if err := sender.Send(m); err != nil {
		c.JSON(badRequest(err.Error()))
		return
	}
	c.AbortWithStatus(http.StatusNoContent)
}

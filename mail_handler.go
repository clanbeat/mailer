package main

import (
	"encoding/json"
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/clanbeat/mailer/src/sender"
	"net/http"
)

func postEmail(c *gin.Context) {
	m := &sender.Mailable{}

	if err := c.BindJSON(m); err != nil {
		c.JSON(badRequest(err.Error()))
		return
	}
	if err := sender.Send(m); err != nil {
		c.JSON(badRequest(err.Error()))
		return
	}
	c.AbortWithStatus(http.StatusNoContent)
}

func getEmail(c *gin.Context) {
	name := c.Param("name")

	m := make(map[string]interface{})
	if !sender.TemplateExists(name) {
		c.JSON(http.StatusNotFound, gin.H{"error_message": "template missing"})
		return
	}
	if testData[name] != nil {
		if err := json.Unmarshal(testData[name], &m); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error_message": "something went wrong"})
			return
		}
	}
	c.HTML(http.StatusOK, name, m)

}

func sendTest(c *gin.Context) {
	name := c.Param("name")
	email := c.Param("email")

	if !sender.TemplateExists(name) {
		c.JSON(http.StatusNotFound, gin.H{"error_message": "template missing"})
		return
	}
	d := make(map[string]interface{})
	if testData[name] != nil {
		if err := json.Unmarshal(testData[name], &d); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error_message": "something went wrong"})
			return
		}
	}
	m := &sender.Mailable{
		To:       email,
		From:     "hello@clanbeat.com",
		Subject:  "Test email",
		Template: name,
		Message:  d,
	}
	if err := sender.Send(m); err != nil {
		c.JSON(badRequest(err.Error()))
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "email sent"})
}

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

	var m map[string]interface{}
	err := json.Unmarshal(testData[name], &m)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.HTML(http.StatusOK, name, m)
}

var testData = map[string][]byte{
	"adminStats":  []byte(`{"user": {"firstName":"Gloria"}}`),
	"betaRequest": []byte(`{"email":"janika.liiv@gmail.com"}`),
}

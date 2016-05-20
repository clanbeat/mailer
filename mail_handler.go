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
	"adminStats": []byte(`{
			"user": {
				"firstName": "Gloria"
			},
			"totals": {
				"highsPercentage": 20,
				"lowsPercentage": 80,
				"lowsCount": 80,
				"highsCount": 20,
				"highsDelta": 20,
				"lowsDelta": -3,
				"publicMoods": 74,
				"privateMoods": 26,
				"publicDelta": 4,
				"privateDelta": -12,
				"moods": 100,
				"moodsDelta": 12
			},
			"start": "1. March",
			"end": "31. March",
			"happiness": [{
				"username": "Janika Liiv",
				"position": 1,
				"percentage": 99,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Gloria Paul",
				"position": 2,
				"percentage": 99,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Janika",
				"position": 3,
				"percentage": 99,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Gloria Paul",
				"position": 4,
				"percentage": 99,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Janika",
				"position": 5,
				"percentage": 99,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Gloria Paul",
				"position": 6,
				"percentage": 99,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			}],
			"activity": [{
				"username": "Janika Liiv",
				"position": 1,
				"count": 10,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Agu Suur",
				"position": 1,
				"count": 5,
				"jobTitle": "Front End Master",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145942953651036900117"
			}],
			"reviews": [{
				"username": "Janika",
				"position": 1,
				"date": "31. April",
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Agu Suur",
				"date": "1. May",
				"position": 2,
				"jobTitle": "Front End Master",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145942953651036900117",
				"link": "tere"
			}]
		}`),
	"betaRequest": []byte(`{"email":"janika.liiv@gmail.com"}`),
}

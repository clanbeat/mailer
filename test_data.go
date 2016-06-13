package main

var testData = map[string][]byte{
	"invitation": []byte(`{
    "inviterName": "Gloria Paul",
    "projectName": "Vikings"
  }`),
	"personalStats": []byte(`{
		"start": "1. March",
		"end": "31. March",
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
		},"reviews": [
			{
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
			}
		]
		}`),
	"adminStats": []byte(`{
			"start": "1. March",
			"end": "31. March",
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
			"happiness": [{
				"username": "Janika Liiv",
				"position": 1,
				"value": 99,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Gloria Paul",
				"position": 2,
				"value": 99,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Janika",
				"position": 3,
				"value": 99,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Gloria Paul",
				"position": 4,
				"value": 99,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Janika",
				"position": 5,
				"value": 99,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Gloria Paul",
				"position": 6,
				"value": 99,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			}],
			"activity": [{
				"username": "Janika Liiv",
				"position": 1,
				"value": 10,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Agu Suur",
				"position": 1,
				"value": 5,
				"jobTitle": "Front End Master",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145942953651036900117"
			}]
		}`),
	"betaRequest": []byte(`{"email":"janika.liiv@gmail.com"}`),
}
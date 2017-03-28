package main

var testData = map[string][]byte{
	"leadPrepare": []byte(`
		{
		  "username": "John Doe",
		  "employeeName": "Karl Suur",
			"goals": [
				{
					"status": 0,
					"content": "finish reading a book",
					"partnerName": "Mari Maasikas",
					"category": "skill"
				},
				{
					"status": 1,
					"content": "Be more fun at the office",
					"partnerName": "Siim Susi",
					"category": "work"
				},
				{
					"status": 0,
					"content": "Be more fun at the office",
					"partnerName": "Siim Susi",
					"category": "manager"
				}
			]
		}
	`),
	"demoRequest": []byte(`
		{
		  "email": "someone@somemail.com",
		  "name": "John Doe",
		  "phoneNumber": "+12313132",
			"company": "Testing Company",
		  "jobTitle": "Team Lead",
		  "teamSize": "2-20",
		  "comments": "some free form text",
		  "location": "London, UK"
		}
	`),
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
			"projectName": "Clanbeat OÜ",
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
			}],
			"reviews": [{
				"username": "Janika Liiv",
				"date": "12. March",
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},
			{
				"username": "Janika Liiv",
				"date": "12. March",
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			}],
			"needReviews": ["Mari Maasikas", "Siim Susi", "Artur Alliksaar", "Marie Under"],
			"oneOnOneLink": "https://beta.clanbeat.com/oneonone"
		}`),
	"betaRequest": []byte(`{"email":"janika.liiv@gmail.com"}`),
	"managerLatestUpdates": []byte(`{
		"user": {
			"firstName": "Janika",
			"lastName": "Liiv",
			"userName": "Janika Liiv",
			"jobTitle": "Back End developer",
			"createdAt": "Sept 2016",
			"picture":"https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518",
			"profileLink": "https://beta.clanbeat.com/profile/18"
		},
		"buttonLink": "https://beta.clanbeat.com/oneonone",
		"moods": [
			{
				"createdAt": "08. November",
				"score": 1,
				"message": "We finally got up some 1-on-1 flow changes, that give me opportunity to urge our users to write more stuff down while preparing. Super."
			},
			{
				"createdAt": "08. November",
				"score": 4,
				"message": "Got email from Rocket Internet (Berlin event followup). This can be a total game changer for Clanbeat. BIG TIME -----  For the demo I believe I will ask for it at the beginning of January 17, We are working now on restructuring the company and scaling up, appraisals and one on one will be at the core of this restructuring."
			},
			{
				"createdAt": "08. November",
				"score": 6,
				"message": "Vilnius Airport is more like pointless bus-station. Flight delays seems rule here."
			}
		]
	}`), "employeeLatestUpdates": []byte(`{
		"user": {
			"firstName": "Janika",
			"lastName": "Liiv",
			"userName": "Janika Liiv",
			"jobTitle": "Back End developer",
			"createdAt": "Sept 2016",
			"picture":"https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518",
			"profileLink": "https://beta.clanbeat.com/profile/18"
		},
		"buttonLink": "https://beta.clanbeat.com/oneonone",
		"moods": [
			{
				"createdAt": "18. November",
				"score": 1,
				"message": "We finally got up some 1-on-1 flow changes, that give me opportunity to urge our users to write more stuff down while preparing. Super."
			},
			{
				"createdAt": "07. November",
				"score": 4,
				"message": "Got email from Rocket Internet (Berlin event followup). This can be a total game changer for Clanbeat. BIG TIME -----  For the demo I believe I will ask for it at the beginning of January 17, We are working now on restructuring the company and scaling up, appraisals and one on one will be at the core of this restructuring."
			},
			{
				"createdAt": "01. November",
				"score": 6,
				"message": "Vilnius Airport is more like pointless bus-station. Flight delays seems rule here."
			}
		]
	}`),
}

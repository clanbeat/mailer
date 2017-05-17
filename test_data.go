package main

var testData = map[string][]byte{
	"oneWeekToReview": []byte(`
		{
		  "username": "John Doe",
		  "partnerName": "Karl Suur",
			"buttonLink": "https://staging.clanbeat.com",
			"scheduledAt": "2. January at 3:04PM",
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
					"category": "employee"
				},
				{
					"status": 0,
					"content": "Be more fun at the office",
					"partnerName": "Siim Susi",
					"category": "manager"
				},
				{
					"status": 1,
					"content": "Clean the table before going to vacation",
					"partnerName": "Siim Susi",
					"category": "employee"
				},
				{
					"status": 1,
					"content": "Find new assistant to Susie, who could do the awesome work Susie needs to be done and has some initiative.",
					"partnerName": "Mart The Manager",
					"category": "employee"
				}
			]
		}
	`),
	"employeePrepare": []byte(`
		{
			"username": "Karl Suur",
			"leadName": "John Doe",
			"buttonLink": "https://staging.clanbeat.com",
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
					"category": "employee"
				},
				{
					"status": 0,
					"content": "Find new assistant to Susie, who could do the awesome work Susie needs to be done and has some initiative.",
					"partnerName": "Siim Susi",
					"category": "manager"
				}
			]
		}
	`),
	"weeklyGoals": []byte(`
		{
			"username": "John Doe",
			"buttonLink": "https://staging.clanbeat.com",
			"projectGoals": [
				{
					"name" : "Team Clanbeaters",
					"goals": [
						{
							"status": 1,
							"content": "Be more fun at the office",
							"partnerName": "Siim Susi",
							"category": "employee"
						},
						{
							"status": 0,
							"content": "Find new assistant to Susie, who could do the awesome work Susie needs to be done and has some initiative.",
							"partnerName": "Siim Susi",
							"category": "manager"
						},
						{
							"status": 0,
							"content": "finish reading a book",
							"partnerName": "Mari Maasikas",
							"category": "skill"
						}
					]
				},
				{
					"name" : "Wafflehouse",
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
							"category": "employee"
						},
						{
							"status": 0,
							"content": "Find new assistant to Susie.",
							"partnerName": "Siim Susi",
							"category": "manager"
						}
					]
				}
			]
		}
	`),
	"leadPrepare": []byte(`
		{
		  "username": "John Doe",
		  "employeeName": "Karl Suur",
			"buttonLink": "https://staging.clanbeat.com",
			"goals": [
				{
					"status": 0,
					"content": "finish reading a book",
					"partnerName": "Mari Maasikas",
					"category": "skill"
				},
				{
					"status": 1,
					"content": "Find new assistant to Susie, who could do the awesome work Susie needs to be done and has some initiative.",
					"partnerName": "Siim Susi",
					"category": "employee"
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
	"leadRequest": []byte(`
		{
		  "username": "John Doe",
		  "employeeName": "Karl Suur",
			"buttonLink": "https://staging.clanbeat.com",
			"scheduledAt": "2. January at 3:04PM",
			"invitationLink": "http://clanbeat.com"
		}
	`),
	"employeeRequest": []byte(`
		{
		  "username": "John Doe",
		  "leadName": "Karl Suur",
			"buttonLink": "https://staging.clanbeat.com",
			"scheduledAt": "2. January at 3:04PM",
			"invitationLink": ""
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
    "projectName": "Vikings",
		"reviewData": {
			"scheduledAt": "2. January at 3:04PM",
			"projectId": 12,
			"employeeId": 1,
			"allDay": false
		}
  }`),
	"invitationToManager": []byte(`{
    "inviterName": "Gloria Paul",
    "projectName": "Vikings",
		"userRole": "admin",
		"reviewData": {
			"scheduledAt": "2. January at 3:04PM",
			"projectId": 12,
			"employeeId": 1,
			"allDay": false
		}
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
				"date": "12. March 12:00PM",
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			}],
			"needReviews": ["Mari Maasikas", "Siim Susi", "Artur Alliksaar", "Marie Under"],
			"oneOnOneLink": "https://beta.clanbeat.com"
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
		"buttonLink": "https://beta.clanbeat.com",
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
		"buttonLink": "https://beta.clanbeat.com",
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

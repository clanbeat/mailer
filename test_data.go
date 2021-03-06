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
					"category": "manager",
					"totalSubgoals": 3,
					"doneSubgoals": 1,
					"subgoals": [{
						"status": 0,
						"content": "start reading the book about having fun",
						"partnerName": "Mari Maasikas",
						"category": "skill"
					},{
						"status": 0,
						"content": "practice smiling",
						"partnerName": "Mari Maasikas",
						"category": "skill"
					},{
						"status": 1,
						"content": "try to make some jokes when talking to someone. They can be awkward in the beginning, but that is ok. It is always a bit awkward, but it will get better over time.",
						"partnerName": "Mari Maasikas",
						"category": "skill"
					}]
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
					"category": "manager",
					"subgoals": [{
						"status": 0,
						"content": "start reading the book about having fun",
						"partnerName": "Mari Maasikas",
						"category": "skill"
					},{
						"status": 0,
						"content": "practice smiling",
						"partnerName": "Mari Maasikas",
						"category": "skill"
					},{
						"status": 1,
						"content": "try to make some jokes when talking to someone. They can be awkward in the beginning, but that is ok. It is always a bit awkward, but it will get better over time.",
						"partnerName": "Mari Maasikas",
						"category": "skill"
					}]
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
							"category": "employee",
							"totalSubgoals": 3,
							"doneSubgoals": 1,
							"subgoals": [{
								"status": 0,
								"content": "start reading the book about having fun",
								"partnerName": "Mari Maasikas",
								"category": "skill"
							},{
								"status": 0,
								"content": "practice smiling",
								"partnerName": "Mari Maasikas",
								"category": "skill"
							},{
								"status": 1,
								"content": "try to make some jokes when talking to someone. They can be awkward in the beginning, but that is ok. It is always a bit awkward, but it will get better over time.",
								"partnerName": "Mari Maasikas",
								"category": "skill"
							}]
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
							"category": "skill",
							"totalSubgoals": 3,
							"doneSubgoals": 2,
							"subgoals": [{
								"status": 1,
								"content": "start reading the book",
								"partnerName": "Mari Maasikas",
								"category": "skill"
							},{
								"status": 1,
								"content": "turn pages in the book and read",
								"partnerName": "Mari Maasikas",
								"category": "skill"
							},{
								"status": 0,
								"content": "think about what I learned from this book when it is finsihed",
								"partnerName": "Mari Maasikas",
								"category": "skill"
							}]
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
							"category": "skill",
							"totalSubgoals": 1,
							"doneSubgoals": 1,
							"subgoals": [{
								"status": 1,
								"content": "try to make some jokes when talking to someone. They can be awkward in the beginning, but that is ok. It is always a bit awkward, but it will get better over time.",
								"partnerName": "Mari Maasikas",
								"category": "skill"
							}]
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
					"category": "skill",
					"subgoals": []
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
	"userMessage": []byte(`
		{
			"toFirstName": "Ivan",
			"fromUsername": "John Doe",
			"fromJobTitle": "Master of Disaster",
			"fromPicture": "https://placekitten.com/200/200",
			"message": "A random message\n\r\n\rEmail me at void@dev.null"
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
				"oldValue": 20,
				"jobTitle": "Developer"
			},{
				"username": "Gloria Paul",
				"position": 2,
				"value": 99,
				"oldValue": 100,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Janika",
				"position": 3,
				"value": 99,
				"oldValue": 20,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Gloria Paul",
				"position": 4,
				"value": 20,
				"oldValue": 20,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Janika",
				"position": 5,
				"value": 56,
				"oldValue": 100,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Gloria Paul",
				"position": 6,
				"value": 99,
				"oldValue": 20,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			}],
			"activity": [{
				"username": "Janika Liiv",
				"position": 1,
				"value": 10,
				"oldValue": 2,
				"jobTitle": "Developer",
				"picture": "https://clanbeat-avatars.s3.eu-central-1.amazonaws.com/145192266114299914518"
			},{
				"username": "Agu Suur",
				"position": 1,
				"value": 5,
				"oldValue": 20,
				"jobTitle": "Front End Master"
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
}

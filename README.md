#Mailer
Sends out emails

connections
  * RabbitMQ
  * sendgrid

## Environment variables

  * ENV - environment name (development/staging/production)
  * PORT - server port
  * SENTRY_DSN - error reporting
  * CLOUDAMQP_URL - RabbitMq uri
  * SENDGRID_USERNAME - username for using sendgrid
  * SENDGRID_PASSWORD - secret for using sendgrid

#Templates
Links to email templates

##Statistics

 * [admin statistics email](http://clanbeat-mailer-staging.herokuapp.com/emails/adminStats)
 * [personal statistics email](http://clanbeat-mailer-staging.herokuapp.com/emails/personalStats)

##Other
* [Latest updates to Employee](http://clanbeat-mailer-staging.herokuapp.com/emails/employeeLatestUpdates)
 * [Latest updates to Manager](http://clanbeat-mailer-staging.herokuapp.com/emails/managerLatestUpdates)
 * [Email log in link](http://clanbeat-mailer-staging.herokuapp.com/emails/loginLink)
 * [User has been deactivated](http://clanbeat-mailer-staging.herokuapp.com/emails/deactivated)
 * [Employee needs to prepare](http://clanbeat-mailer-staging.herokuapp.com/emails/employeePrepare)
 * [Manager requested a 1-on-1 with employee](http://clanbeat-mailer-staging.herokuapp.com/emails/employeeRequest)
 * [First login to clan creator](http://clanbeat-mailer-staging.herokuapp.com/emails/firstLogin)
 * [Invitation](http://clanbeat-mailer-staging.herokuapp.com/emails/invitation)
 * [Invitation Accepted](http://clanbeat-mailer-staging.herokuapp.com/emails/invitationAccepted)
 * [Invitation Pending](http://clanbeat-mailer-staging.herokuapp.com/emails/invitationPending)
 * [Manager needs to prepare](http://clanbeat-mailer-staging.herokuapp.com/emails/leadPrepare)
 * [Employee requested for a 1-on-1 with manager](http://clanbeat-mailer-staging.herokuapp.com/emails/leadRequest)
 * [No posts for 2 weeks](http://clanbeat-mailer-staging.herokuapp.com/emails/noPosts)
 * [Not subscribed after trial end](http://clanbeat-mailer-staging.herokuapp.com/emails/notSubscribed)
 * [Payment failed](http://clanbeat-mailer-staging.herokuapp.com/emails/paymentFailed)
 * [Daily reminder](http://clanbeat-mailer-staging.herokuapp.com/emails/reminder)
 * [Trial ends in a week](http://clanbeat-mailer-staging.herokuapp.com/emails/trialEnding)
 * [Unsubscribed](http://clanbeat-mailer-staging.herokuapp.com/emails/unsubscribed)
 * [User upgraded to Admin](http://clanbeat-mailer-staging.herokuapp.com/emails/upgradedToAdmin)
 * [User upgraded to Manager](http://clanbeat-mailer-staging.herokuapp.com/emails/upgradedToManager)

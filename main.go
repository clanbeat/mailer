package main

import (
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/clanbeat/broker"
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/clanbeat/errortracker"
	"github.com/clanbeat/mailer/src/config"
	"github.com/clanbeat/mailer/src/sender"
	"log"
	"os"
	"path/filepath"
)

var errorTracker *errortracker.Tracker
var env *config.AppConfig
var brokerConn *broker.Connection

func main() {

	//set up environment
	templatePath := templatePath()
	env = config.New(os.Getenv("ENV"))

	//set up error reporting
	var err error
	errorTracker, err = errortracker.New(os.Getenv("SENTRY_DSN"), env.Env)
	if err != nil {
		log.Fatal(err)
	}

	//set up mailer
	sender.SendgridSetup(os.Getenv("SENDGRID_USERNAME"), os.Getenv("SENDGRID_PASSWORD"))
	if err := sender.CacheTemplates(templatePath); err != nil {
		fatalError(err)
	}

	//set up queue handling
	brokerConn, err = broker.New(os.Getenv("CLOUDAMQP_URL"), errorTrackerFunc())
	if err != nil {
		fatalError(err)
	}
	defer brokerConn.Close()
	registerBrokerHandlers()

	//start server
	initRoutes(templatePath)

	//don't close before errors have been reported
	errorTracker.Wait()
}

func templatePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fatalError(err)
	}
	return dir + "/templates"
}

func fatalError(err error) {
	if err != nil {
		errorTracker.ErrorAndWait(err)
		log.Fatal(err)
	}
}

func errorTrackerFunc() func(err error) {
	return func(err error) {
		errorTracker.Error(err)
	}
}

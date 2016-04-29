package main

import (
	"fmt"
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/clanbeat/errortracker"
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/clanbeat/mailer/src/config"
	"github.com/clanbeat/mailer/src/sender"
	"log"
	"os"
	"path/filepath"
)

var router *gin.Engine
var errorTracker *errortracker.Tracker
var env *config.AppConfig

func main() {

	//set up environment
	env = config.New(os.Getenv("ENV"))

	//set up error reporting
	var err error
	errorTracker, err = errortracker.New(os.Getenv("SENTRY_DSN"), env.Env)
	if err != nil {
		log.Fatal(err)
	}

	//set up mailer
	sender.SendgridSetup(os.Getenv("SENDGRID_USERNAME"), os.Getenv("SENDGRID_PASSWORD"))
	if err := sender.CacheTemplates(templatePath()); err != nil {
		fatalError(err)
	}

	//set up routes
	initRoutes()

	//start server
	log.Println(fmt.Sprintf("=> Starting in %s on http://0.0.0.0:%s", env.Env, os.Getenv("PORT")))
	router.Run(":" + os.Getenv("PORT"))

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

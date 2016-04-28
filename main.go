package main

import (
	"fmt"
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/clanbeat/errortracker"
	"github.com/clanbeat/mailer/Godeps/_workspace/src/github.com/gin-gonic/gin"
	"github.com/clanbeat/mailer/src/config"
	"log"
	"os"
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

	//set up routes
	initRoutes()

	//start server
	log.Println(fmt.Sprintf("=> Starting in %s on http://0.0.0.0:%s", env.Env, os.Getenv("PORT")))
	router.Run(":" + os.Getenv("PORT"))

	//don't close before errors have been reported
	errorTracker.Wait()
}

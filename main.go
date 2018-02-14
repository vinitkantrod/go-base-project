package main

import (
	"io/ioutil"
	"os"

	"github.com/vinitkantrod/go-base-project/config"
	"github.com/vinitkantrod/go-base-project/database"
	"github.com/vinitkantrod/go-base-project/logging"
)

func getEnv() string {
	env := os.Getenv("GO_CAMPAIGNS")
	if env != "" {
		logging.Info.Println("Got Environment : ", env)
	} else {
		os.Setenv("GO_CAMPAIGNS", "production")
		env = os.Getenv("GO_CAMPAIGNS")
	}
	return env
}

func main() {
	logging.Init(ioutil.Discard)
	logging.Info.Println("Initialized Logging Module...")
	env := getEnv()
	config.Init(env)
	database.Init(config.Config.MongoHost, config.Config.MongoDbName)
}

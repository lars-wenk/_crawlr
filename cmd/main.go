package main

import (
	"github.com/joho/godotenv"
	"github.com/lars-wenk/_crawlr/internal/config"
	log "github.com/sirupsen/logrus"
)

var conf config.Config

func init() {
	var err error

	// load .env file
	_ = godotenv.Load()

	// init config
	conf, err = config.NewConfig()
	if err != nil {
		log.Panic(err)
	}

}

func main() {
	crawlr := crawlr.NewController(conf)
	crawlr.Start()
}

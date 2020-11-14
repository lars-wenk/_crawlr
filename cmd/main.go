package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/lars-wenk/_crawlr/internal/config"
	crawlr "github.com/lars-wenk/_crawlr/internal/crawler"
	"github.com/lars-wenk/_crawlr/pkg/logs"
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

	// configure logging
	withColors := false

	if conf.AppEnv == "develop" {
		withColors = true
	}

	logLevel, err := log.ParseLevel(conf.LogLevel)
	if err != nil {
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)

	log.SetFormatter(&logs.LogFormatter{
		WithColors: withColors,
	})
	log.StandardLogger().SetReportCaller(true)

}

func main() {
	ctx, cancelMainContext := context.WithCancel(context.Background())
	log.Info("lets go")

	crawlr := crawlr.NewController(conf, ctx)
	crawlr.Start()
	cancelMainContext()

}

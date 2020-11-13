package crawlr

import (
	"time"

	"github.com/lars-wenk/_crawlr/internal/broker"
	"github.com/lars-wenk/_crawlr/internal/config"
)

const (
	crawlTimeout  = 30 * time.Second
	updateTimeout = 60 * time.Second
	cleanTimeout  = 60 * time.Second
)

type Controller interface {
	Start()
}

type controller struct {
	conf config.Config
}

func NewController(conf config.Config) Controller {
	return &controller{
		conf: conf,
	}
}

func (c *controller) Start() {
	c.crawlComdirect(c.conf)
}

func (c *controller) crawlComdirect(conf config.Config) {
	cc := broker.NewComdirectCrawler(c.conf)
	cc.NewComdirectCrawler()

	return
}

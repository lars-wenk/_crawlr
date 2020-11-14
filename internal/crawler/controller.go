package crawlr

import (
	"context"
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
	ctx  context.Context
}

func NewController(conf config.Config, ctx context.Context) Controller {
	return &controller{
		conf: conf,
		ctx:  ctx,
	}
}

func (c *controller) Start() {
	c.crawlComdirect(c.conf)
}

func (c *controller) crawlComdirect(conf config.Config) {
	cc := broker.NewComdirectCrawler(c.conf)
	cc.GetAuth()

	return
}

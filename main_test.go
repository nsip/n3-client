package main

import (
	"testing"

	c "./config"
	q "./query"
	s "./send"
	w "./send/filewatcher"
	r "./send/rest"
)

func TestMain(t *testing.T) {
	cfg := c.GetConfig("./config.toml", "./config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog, true) }()
	s.Init(cfg)
	q.Init(cfg)

	done := make(chan string)
	go r.HostHTTPForPubAsync()
	go w.StartFileWatcherAsync()
	<-done
}

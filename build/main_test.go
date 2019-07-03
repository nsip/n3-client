package main

import (
	"testing"

	c "../config"
	fw "../filewatcher"
	g "../global"
	pub "../publish"
	"../query"
	"../rest"
	u "github.com/cdutwhu/go-util"
)

var (
	ph = u.PanicHandle
)

func TestMain(t *testing.T) {
	g.Cfg = c.FromFile("../build/config.toml")
	defer func() { ph(recover(), g.Cfg.ErrLog) }()
	pub.InitClient(g.Cfg)
	query.InitClient(g.Cfg)
	rest.InitClient(g.Cfg)

	done := make(chan string)
	go rest.HostHTTPAsync()
	go fw.StartFileWatcherAsync()
	<-done
}

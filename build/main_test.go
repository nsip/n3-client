package main

import (
	"testing"

	fw "../filewatcher"
	g "../global"
	"../rest"
	u "github.com/cdutwhu/go-util"
)

var (
	ph = u.PanicHandle
)

func TestMain(t *testing.T) {
	g.Init()
	defer func() { ph(recover(), g.Cfg.ErrLog) }()

	done := make(chan string)
	go rest.HostHTTPAsync()
	go fw.StartFileWatcherAsync()
	<-done
}

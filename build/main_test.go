package main

import (
	"testing"

	u "github.com/cdutwhu/go-util"
	fw "github.com/nsip/n3-client/filewatcher"
	g "github.com/nsip/n3-client/global"
	"github.com/nsip/n3-client/rest"
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

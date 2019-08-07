package main

import (
	fw "github.com/nsip/n3-client/filewatcher"
	g "github.com/nsip/n3-client/global"
	"github.com/nsip/n3-client/rest"

	u "github.com/cdutwhu/go-util"
)

var (
	ph = u.PanicHandle
)

func main() {
	g.Init()
	defer func() { ph(recover(), g.Cfg.ErrLog) }()

	done := make(chan string)
	go rest.HostHTTPAsync()
	go fw.StartFileWatcherAsync()
	<-done
}

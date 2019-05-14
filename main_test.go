package main

import (
	"testing"

	cfg "./config"
	"./query"
	"./rest"
	"./send"
	w "./send/filewatcher"
)
 
func TestMain(t *testing.T) {
	CFG := cfg.FromFile("./config.toml", "./config/config.toml")
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	send.InitClient(CFG)
	query.InitClient(CFG)
	rest.InitClient(CFG)

	done := make(chan string)
	go rest.HostHTTPAsync()
	go w.StartFileWatcherAsync()
	<-done
}

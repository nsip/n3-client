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
	send.InitFrom(CFG)
	query.InitFrom(CFG)
	rest.InitFrom(CFG)

	done := make(chan string)
	go rest.HostHTTPAsync()
	go w.StartFileWatcherAsync()
	<-done
}

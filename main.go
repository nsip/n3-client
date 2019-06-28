package main

import (
	cfg "./config"
	"./query"
	"./rest"
	pub "./publish"
	fw "./publish/filewatcher"
)

func main() {
	CFG := cfg.FromFile("./config.toml", "./config/config.toml")
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	pub.InitClient(CFG)
	query.InitClient(CFG)
	rest.InitClient(CFG)

	done := make(chan string)
	go rest.HostHTTPAsync()
	go fw.StartFileWatcherAsync()
	<-done
}

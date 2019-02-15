package main

// func SendXmlToDataStore(filename string) {
// 	fi, err := os.Lstat(filename)
// }

import (
	c "./config"
	q "./query"
	s "./send"
	w "./send/filewatcher"
	r "./send/rest"
)

func main() {
	cfg := c.GetConfig("./config.toml", "./config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog) }()
	s.Init(cfg)
	q.Init(cfg)

	done := make(chan string)
	go r.HostHTTPForPubAsync()
	go w.StartFileWatcherAsync()
	<-done
}

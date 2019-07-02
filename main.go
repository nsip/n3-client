package main

import (
	"fmt"
	"strings"

	cfg "./config"
	fw "./filewatcher"
	pub "./publish"
	"./query"
	"./rest"

	u "github.com/cdutwhu/go-util"
)

var (
	PE  = u.PanicOnError
	PE1 = u.PanicOnError1
	PH  = u.PanicHandle
	PC  = u.PanicOnCondition

	sFF = strings.FieldsFunc
	sC  = strings.Contains
	sJ  = strings.Join

	fPln = fmt.Println
	fPf  = fmt.Printf
	fEf  = fmt.Errorf
	fSf  = fmt.Sprintf
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

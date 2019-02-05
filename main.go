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

// func SendXmlToDataStore(filename string) {
// 	defer PH("", false)
// 	// defer PHE("", false, func(emsg string, params ...interface{}) {
// 	// 	fmt.Println(emsg)
// 	// 	fmt.Println(params[0])
// 	// }, "do more things?")

// 	fi, err := os.Lstat(filename)
// 	PE(err)
// 	PC(fi.Mode().IsDir(), epf("%s is a directory", filename))
// 	PC(!sHS(filename, ".xml"), epf("%s is not an XML file", filename))

// 	file, err := os.Open(filename)
// 	PE1(err, fSf("Cannot read in file %s\n", filename))
// }

func main() {
	cfg := c.GetConfig("./config.toml", "./config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog, true) }()
	s.Init(cfg)
	q.Init(cfg)

	done := make(chan string)
	go r.HostHTTPForPubAsync()
	go w.StartFileWatcherAsync()
	<-done
}

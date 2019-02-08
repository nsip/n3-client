package send

import (
	"io/ioutil"
	"testing"
	"time"

	c "../config"
)

func TestJunk(t *testing.T) {
	defer func() { PH(recover(), Cfg.Global.ErrLog) }()
	TestN3LoadConfig(t)
	Junk(10)
	time.Sleep(2 * time.Second)
}

/************************************************************/

func TestN3LoadConfig(t *testing.T) {
	cfg := c.GetConfig("./config.toml", "../config/config.toml")
	// fPln(cfg.Grpc)
	// fPln(cfg.Filewatcher)
	// fPln(cfg.Path)
	Init(cfg)
}

func TestSendSif(t *testing.T) {
	defer func() { PH(recover(), Cfg.Global.ErrLog) }()
	TestN3LoadConfig(t)

	// xmlfile := "../inbound/sif/staffpersonal.xml"
	xmlfile := "../inbound/sif/nswdig.xml"
	bytes := Must(ioutil.ReadFile(xmlfile)).([]byte)
	Sif(string(bytes))
}

func TestSendXapi(t *testing.T) {
	defer func() { PH(recover(), Cfg.Global.ErrLog) }()
	TestN3LoadConfig(t)

	jsonfile := "../inbound/xapi/xapifile.json"
	bytes := Must(ioutil.ReadFile(jsonfile)).([]byte)
	Xapi(string(bytes))
}

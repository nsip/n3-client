package send

import (
	"io/ioutil"
	"testing"
	"time"

	c "../config"
)

func TestJunk(t *testing.T) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	TestN3LoadConfig(t)
	Junk(10)
	time.Sleep(2 * time.Second)
}

/************************************************************/

func TestN3LoadConfig(t *testing.T) {
	InitFrom(c.FromFile("./config.toml", "../config/config.toml"))
}

func TestSendSif(t *testing.T) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	TestN3LoadConfig(t)

	xmlfile := "../inbound/sif/sif.xml"	
	bytes := Must(ioutil.ReadFile(xmlfile)).([]byte)
	Sif(string(bytes))
}

func TestSendXapi(t *testing.T) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	TestN3LoadConfig(t)

	jsonfile := "../inbound/xapi/xapi.json"	
	bytes := Must(ioutil.ReadFile(jsonfile)).([]byte)
	Xapi(string(bytes))
}

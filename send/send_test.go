package send

import (
	"io/ioutil"
	"testing"
	"time"

	c "../config"
	g "../global"
)

func TestJunk(t *testing.T) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	TestN3LoadConfig(t)
	Junk(10)
	time.Sleep(2 * time.Second)
}

/************************************************************/

func TestN3LoadConfig(t *testing.T) {
	InitClientFrom(c.FromFile("./config.toml", "../config/config.toml"))
}

func TestToNode(t *testing.T) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	TestN3LoadConfig(t)

	datafile := "../inbound/xapi/xapiC.json" //              *** change file ***
	bytes := Must(ioutil.ReadFile(datafile)).([]byte)
	ToNode(string(bytes), g.XAPI) //                         *** change data type ***
	time.Sleep(5 * time.Second)
}

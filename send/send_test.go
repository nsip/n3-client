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
	InitClient(c.FromFile("./config.toml", "../config/config.toml"))
}

func TestToNode(t *testing.T) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	TestN3LoadConfig(t)

	datafile := "../inbound/xapi/xapiC.json" //                        *** change file ***
	bytes := Must(ioutil.ReadFile(datafile)).([]byte)
	IDs, _, _, _ := ToNode(string(bytes), "id", "xapi") //             *** change idmark ***
	time.Sleep(1 * time.Second)
	for _, id := range IDs {
		fPln("sent:", id)
	}
}

package send

import (
	"io/ioutil"
	"testing"
	"time"

	c "../config"
	pp "../preprocess"
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

	file := "../inbound/hsie/history/stage5/overview.json" //           *** change <file> ***
	data := string(Must(ioutil.ReadFile(file)).([]byte))
	data = pp.FmtJSONStr(data, "../preprocess/util/", "./") //           *** format json string ***
	IDs, _, _, _ := ToNode(data, "id", "hsie")              //           *** change <idmark>, <default root> ***
	time.Sleep(1 * time.Second)
	for _, id := range IDs {
		fPln("sent:", id)
	}
}

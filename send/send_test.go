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

	datafile := "../inbound/hsie/geography/stage3/content.json" //             *** change file ***
	data := string(Must(ioutil.ReadFile(datafile)).([]byte))

	// ***
	// data = u.FmtJSONStr(data, "../util/")
	// if ok := checkColon(data); !ok {
	// 	t.Errorf("%s has colon in field value, fix them before sending", datafile)
	// 	return
	// }
	// if ok := checkHyphen(data); !ok {
	// 	t.Errorf("%s has hyphen in field value, fix them before sending", datafile)
	// 	return
	// }
	// ***

	IDs, _, _, _ := ToNode(data, "id", "hsie") //                        *** change <idmark>, <default root> ***
	time.Sleep(1 * time.Second)
	for _, id := range IDs {
		fPln("sent:", id)
	}
}

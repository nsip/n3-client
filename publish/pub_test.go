package publish

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

	// for i := 1; i <= 5; i++ {
	// 	file := fSf("../inbound/hsie/geography/stage%d/overview.json", i) // *** change <file> ***
	// 	json := string(Must(ioutil.ReadFile(file)).([]byte))
	// 	IDs, _, _, _, _ := Pub2Node(json, "id", "Overview") //             *** change <idmark> <dfltRoot> ***
	// 	time.Sleep(1 * time.Second)
	// 	for _, id := range IDs {
	// 		fPln("sent:", id)
	// 	}
	// }

	file := "../inbound/xapi/xapi.json" //                  *** change <file> ***
	json := string(Must(ioutil.ReadFile(file)).([]byte))
	IDs, _, _, _, _ := Pub2Node(json, "id", "xapi") //      *** change <idmark> <dfltRoot> ***
	time.Sleep(1 * time.Second)
	for _, id := range IDs {
		fPln("sent:", id)
	}
}

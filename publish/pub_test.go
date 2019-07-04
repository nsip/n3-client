package publish

import (
	"io/ioutil"
	"testing"
	"time"

	c "../config"
	g "../global"
)

func TestJunk(t *testing.T) {
	defer func() { ph(recover(), g.Cfg.ErrLog) }()
	TestN3LoadConfig(t)

	g.CurCtx = g.Cfg.RPC.CtxList[0]
	Junk(g.CurCtx, 3)
	time.Sleep(200 * time.Millisecond)
}

/************************************************************/

func TestN3LoadConfig(t *testing.T) {
	InitClient(c.FromFile("../build/config.toml"))
}

func TestToNode(t *testing.T) {
	defer func() { ph(recover(), g.Cfg.ErrLog) }()
	TestN3LoadConfig(t)

	g.CurCtx = g.Cfg.RPC.CtxList[0]

	for i := 1; i <= 5; i++ {
		file := fSf("../inbound/hsie/history/stage%d/overview.json", i) // *** change <file> ***
		json := string(must(ioutil.ReadFile(file)).([]byte))
		IDs, _, _, _, _ := Pub2Node(g.CurCtx, json, "id", "Overview") //     *** change <idmark> <dfltRoot> ***
		time.Sleep(1 * time.Second)
		for _, id := range IDs {
			fPln("sent:", id)
		}
	}

	// file := "../inbound/xapi/xapi.json" //                  *** change <file> ***
	// json := string(must(ioutil.ReadFile(file)).([]byte))
	// IDs, _, _, _, _ := Pub2Node(g.CurCtx, json, "id", "xapi") //      *** change <idmark> <dfltRoot> ***
	// time.Sleep(1 * time.Second)
	// for _, id := range IDs {
	// 	fPln("sent:", id)
	// }
}

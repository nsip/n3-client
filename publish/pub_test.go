package publish

import (
	"io/ioutil"
	"testing"
	"time"

	c "../config"
	g "../global"
)

func TestN3LoadConfig(t *testing.T) {
	InitClient(c.FromFile("../build/config.toml"))
}

func TestJunk(t *testing.T) {
	defer func() { ph(recover(), g.Cfg.ErrLog) }()
	TestN3LoadConfig(t)

	g.CurCtx = g.Cfg.RPC.CtxList[0]
	Junk(g.CurCtx, 3)
	time.Sleep(200 * time.Millisecond)
}

func TestToNode(t *testing.T) {
	defer func() { ph(recover(), g.Cfg.ErrLog) }()
	TestN3LoadConfig(t)

	g.CurCtx = g.Cfg.RPC.CtxList[0]

	// for i := 1; i <= 5; i++ {
	// 	file := fSf("../inbound/hsie/history/stage%d/overview.json", i) // *** change <file> ***
	// 	json := string(must(ioutil.ReadFile(file)).([]byte))
	// 	IDs, _, _, _, _ := Pub2Node(g.CurCtx, json, "id", "Overview") //     *** change <idmark> <dfltRoot> ***
	// 	time.Sleep(1 * time.Second)
	// 	for _, id := range IDs {
	// 		fPln("sent:", id)
	// 	}
	// }

	file := "../inbound/xapi/xapi.json" //                  *** change <file> ***
	json := string(must(ioutil.ReadFile(file)).([]byte))
	IDs, _, _, _, _ := Pub2Node(g.CurCtx, json, "id", "xapi") //      *** change <idmark> <dfltRoot> ***
	time.Sleep(1 * time.Second)
	for _, id := range IDs {
		fPln("sent:", id)
	}
}

func TestPrictrlToNode(t *testing.T) {
	defer func() { ph(recover(), g.Cfg.ErrLog) }()
	TestN3LoadConfig(t)

	lCtxList := len(g.Cfg.RPC.CtxList)
	g.CurCtx = g.Cfg.RPC.CtxList[lCtxList-2] //                   *** use THE LAST BUT ONE for <privacy control> measurement ***

	fPln(g.CurCtx)

	file := "../inbound/prictrl/xapi.json" //                     *** change <file> ***
	json := string(must(ioutil.ReadFile(file)).([]byte))
	IDs, _, _, _, _ := Pub2Node(g.CurCtx, json, "id", "xapi") //  *** change <idmark> <dfltRoot> ***
	time.Sleep(1 * time.Second)
	for _, id := range IDs {
		fPln("sent:", id)
	}

}

func TestCtxidToNode(t *testing.T) {
	defer func() { ph(recover(), g.Cfg.ErrLog) }()
	TestN3LoadConfig(t)

	lCtxList := len(g.Cfg.RPC.CtxList)
	g.CurCtx = g.Cfg.RPC.CtxList[lCtxList-1] //                   *** use THE LAST for <context-id> measurement ***

	fPln(g.CurCtx)

	Send(g.CurCtx, "xapi11", "comment 22", "4947ED1F-1E94-4850-8B8F-35C653F51E9G") // *** ctx, id, comment ***
	time.Sleep(1 * time.Second)
}

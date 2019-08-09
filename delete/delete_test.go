package delete

import (
	"testing"
	"time"

	g "github.com/nsip/n3-client/global"
)

func TestDelete(t *testing.T) {
	g.Init()
	defer func() { ph(recover(), g.Cfg.ErrLog) }()

	CurCtx := g.Cfg.RPC.CtxList[0]

	DelBat(CurCtx, "5e8d383b-13d2-481f-8db5-c16376279566", "A759FF45-4ABD-4A59-B31B-BB0D3CA66ADC", "738F4DF5-949F-4380-8186-8252440A6F6F")
	time.Sleep(1 * time.Second)
}

package delete

import (
	"testing"
	"time"

	g "../global"
)

func TestDelete(t *testing.T) {
	g.Init()
	defer func() { ph(recover(), g.Cfg.ErrLog) }()

	DelBat(g.CurCtx, "46F3503C-C6D5-4BE8-ABB8-D3B5FE853948", "A759FF45-4ABD-4A59-B31B-BB0D3CA66ADC", "738F4DF5-949F-4380-8186-8252440A6F6F")
	time.Sleep(1 * time.Second)
}

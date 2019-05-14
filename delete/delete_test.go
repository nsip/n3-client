package delete

import (
	"testing"
	"time"

	c "../config"
)

func TestN3LoadConfig(t *testing.T) {
	Init(c.FromFile("./config.toml", "../config/config.toml"))
}

func TestDelete(t *testing.T) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	TestN3LoadConfig(t)
	Del("03edb5af-b593-41b4-a4cc-6f34bec90408")
	time.Sleep(1 * time.Second)
}

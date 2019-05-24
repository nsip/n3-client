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
	DelBat("46F3503C-C6D5-4BE8-ABB8-D3B5FE853948", "A759FF45-4ABD-4A59-B31B-BB0D3CA66ADC", "738F4DF5-949F-4380-8186-8252440A6F6F")
	time.Sleep(1 * time.Second)
}

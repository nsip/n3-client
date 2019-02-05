package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	cfg := GetConfig("./config.toml")
	fPln(cfg.RPC)
	fPln(cfg.Filewatcher)
	fPln(cfg.Global.ErrLog)
}

func TestSave(t *testing.T) {
	cfg := GetConfig("./config.toml")
	cfg.Save()

	cfg1 := GetConfig("./config.toml")
	fPln(cfg1.RPC)
	fPln(cfg1.Filewatcher)
}

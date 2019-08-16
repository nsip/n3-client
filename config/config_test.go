package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	cfg := FromFile("../config.toml")
	fPln(cfg.Path)
	fPln(cfg.ErrLog)
	fPln(cfg.FileWatcher)
	fPln(cfg.WebService)
	fPln(cfg.Group)
	fPln(cfg.Route)
	fPln(cfg.RPC)
	fPln(cfg.Query)
	fPln(cfg.Debug)
}

// func TestSave(t *testing.T) {
// 	cfg := FromFile("../config.toml")
// 	cfg.WebService.VerMajor = 1
// 	cfg.Save()
// 	cfg1 := FromFile("../config.toml")
// 	fPln(cfg1.RPC)
// 	fPln(cfg1.FileWatcher)
// }

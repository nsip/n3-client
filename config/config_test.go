package config

import (
	"testing"
)

// 1) Ctrl+, -> Settings, 2) look for "go test", 3) [Build Flags], 4) [Add Item] add "-v"

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

package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	cfg := FromFile("../build/config.toml")
	fPln(cfg.Rest.PathGQL)
	fPln(cfg.RPC)
	fPln(cfg.Filewatcher)
	fPln(cfg.ErrLog)
	fPln(cfg.RPC.CtxList)
	fPln(cfg.RPC.CtxPrivDef)
	fPln(cfg.RPC.CtxPrivID)
	fPln(cfg.Query.ParamPathDir)
}

func TestSave(t *testing.T) {
	cfg := FromFile("../build/config.toml")
	cfg.Save()

	cfg1 := FromFile("../build/config.toml")
	fPln(cfg1.RPC)
	fPln(cfg1.Filewatcher)
}

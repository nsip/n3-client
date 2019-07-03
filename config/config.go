package config

import (
	"os"

	"github.com/burntsushi/toml"
)

type filewatcher struct {
	Dir string
}

type rest struct {
	Port       int
	PathTest   string
	PathID     string
	PathPub    string
	PathGQL    string
	PathDel    string
	PathObj    string
	PathScm    string
	PathGQLTxt string
}

type rpc struct {
	Namespace string
	Server    string
	Port      int
	CtxList   []string
}

type query struct {
	SchemaDir    string
	SampleTxtDir string
	ParamPathDir string
}

// Config is toml
type Config struct {
	Path        string
	ErrLog      string
	Filewatcher filewatcher
	Rest        rest
	RPC         rpc
	Query       query
}

// FromFile :
func FromFile(cfgfiles ...string) *Config {
	for _, f := range cfgfiles {
		if _, e := os.Stat(f); e == nil {
			cfg := &Config{Path: f}
			return cfg.set()
		}
	}
	panic("config file error")
}

// set is
func (cfg *Config) set() *Config {
	defer func() { ph(recover(), cfg.ErrLog) }()
	path := cfg.Path /* make a copy of original path for restoring */
	must(toml.DecodeFile(cfg.Path, cfg))
	cfg.Path = path
	return cfg
}

// Save is
func (cfg *Config) Save() {
	defer func() { ph(recover(), cfg.ErrLog) }()
	f := must(os.OpenFile(cfg.Path, os.O_WRONLY|os.O_TRUNC, 0666)).(*os.File)
	defer f.Close()
	pe(toml.NewEncoder(f).Encode(cfg))
}

package config

import (
	"os"

	"github.com/burntsushi/toml"
)

type global struct {
	ErrLog string
}

type filewatcher struct {
	DirSif  string
	DirXapi string
}

type rest struct {
	Port      int
	PathSend string	
	PathGQL   string
}

type rpc struct {
	Namespace string
	Server    string
	Port      int
	Ctx       string
	CtxMeta   string	
}

// Config is toml
type Config struct {
	Path        string
	Global      global
	Filewatcher filewatcher
	Rest        rest
	RPC         rpc
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
	defer func() { PH(recover(), "./log.txt") }()
	path := cfg.Path /* make a copy of original path for restoring */
	Must(toml.DecodeFile(cfg.Path, cfg))
	cfg.Path = path
	return cfg
}

// Save is
func (cfg *Config) Save() {
	defer func() { PH(recover(), cfg.Global.ErrLog) }()
	f := Must(os.OpenFile(cfg.Path, os.O_WRONLY|os.O_TRUNC, 0666)).(*os.File)
	defer f.Close()
	PE(toml.NewEncoder(f).Encode(cfg))
}

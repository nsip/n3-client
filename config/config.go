package config

import (
	"os"
	"reflect"

	"github.com/burntsushi/toml"
)

type filewatcher struct {
	Dir string
}

type webservice struct {
	Port    int
	Version string
}

type group struct {
	APP string
	API string
}

type route struct {
	Greeting string
	ID       string
	Pub      string
	GQL      string
	GQL2     string
	Del      string
	Obj      string
	Scm      string
	GQLTxt   string
	Upload   string
}

type rpc struct {
	Namespace  string
	Server     string
	Port       int
	CtxList    []string
	CtxPrivDef string
	CtxPrivID  string
}

type query struct {
	SchemaDir    string
	SampleTxtDir string
	ParamPathDir string
}

type debug struct {
	TrialPub bool
	TrialQry bool
}

// ********************************************** //

// Config is toml
type Config struct {
	Path        string
	ErrLog      string
	FileWatcher filewatcher
	WebService  webservice
	Group       group
	Route       route
	RPC         rpc
	Query       query
	Debug       debug
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
	return cfg.modCfg()
}

func (cfg *Config) modCfg() *Config {
	// *** replace version ***
	cfgws := cfg.WebService
	ver := fSf("v%s", cfgws.Version)
	v := reflect.ValueOf(cfg.Route)
	for i := 0; i < v.NumField(); i++ {
		vv := S(v.Field(i).Interface().(string)).Replace("v#", ver).V()
		reflect.ValueOf(&cfg.Route).Elem().Field(i).SetString(vv)
	}
	return cfg
}

// Save is
// func (cfg *Config) Save() {
// 	defer func() { ph(recover(), cfg.ErrLog) }()
// 	f := must(os.OpenFile(cfg.Path, os.O_WRONLY|os.O_TRUNC, 0666)).(*os.File)
// 	defer f.Close()
// 	pe(toml.NewEncoder(f).Encode(cfg))
// }

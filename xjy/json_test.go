package xjy

import (
	"io/ioutil"
	"testing"

	c "../config"
)

func TestJSONGetObjID(t *testing.T) {
	cfg := c.GetConfig("./config.toml", "../config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog) }()
	jsonbytes := Must(ioutil.ReadFile("./files/xapifile.json")).([]byte)
	id := JSONGetObjID(string(jsonbytes), "id", " ~ ")
	fPln(id)
}

func TestJSONScanObjects(t *testing.T) {
	cfg := c.GetConfig("./config.toml", "../config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog) }()
	jsonbytes := Must(ioutil.ReadFile("./files/xapifile.json")).([]byte)

	idx := 1
	JSONModelInfo(string(jsonbytes), "id", " ~ ",
		func(p, v string) {
			fPf("%-5d: %-80s:: %s\n", idx, p, v)
			idx++
		},
		func(p, v string, n int) {
			fPf("%-5d: %-80s:: %s -- [%d]\n", idx, p, v, n)
			idx++
		},
	)
}

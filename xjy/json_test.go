package xjy

import (
	"io/ioutil"
	"testing"

	c "../config"
)

func TestJSONGetObjID(t *testing.T) {
	cfg := c.FromFile("./config.toml", "../config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog) }()
	jsonbytes := Must(ioutil.ReadFile("./files/sample.json")).([]byte)
	id, autoID := JSONGetObjID(string(jsonbytes), "id", "DefaultRoot", PATH_DEL)
	fPln(id, autoID)
}

func TestJSONScanObjects(t *testing.T) {
	cfg := c.FromFile("./config.toml", "../config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog) }()

	xapibytes := Must(ioutil.ReadFile("./files/sample.json")).([]byte)
	xapi := Str(xapibytes)
	xapi.SetEnC()

	mapStructRecord := map[string][]string{}
	procIdx := 1
	JSONObjScan(xapi.V(), "id", "ROOT",
		func(p, id string, v []string, lastObjTuple bool) {
			if _, ok := mapStructRecord[p]; !ok {
				mapStructRecord[p] = v
				fPf("S%3d ---> %-70s:: %s\n", procIdx, p, sJ(v, CHILD_DEL))
				procIdx++
			}
		},
		func(p, id string, n int, lastObjTuple bool) {
			fPf("A%3d ---> %-70s[]%s -- [%d]\n", procIdx, p, id, n)
			procIdx++
		},
	)
}

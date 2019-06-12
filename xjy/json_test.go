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
	id, autoID, addRoot := JSONGetObjID(string(jsonbytes), "id", "DefaultRoot", PATH_DEL)
	fPln(id, autoID, addRoot)
}

func TestJSONScanObjects(t *testing.T) {
	cfg := c.FromFile("./config.toml", "../config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog) }()

	data := string(Must(ioutil.ReadFile("./files/content.json")).([]byte))
	mStructRecord := map[string][]string{}
	procIdx := 1
	JSONObjScan(data, "id", "ROOT",
		func(p, id string, v []string, lastObjTuple bool) {
			if _, ok := mStructRecord[p]; !ok {
				mStructRecord[p] = v
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

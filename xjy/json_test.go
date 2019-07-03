package xjy

import (
	"io/ioutil"
	"testing"

	c "../config"
)

func TestJSONGetObjID(t *testing.T) {
	cfg := c.FromFile("../build/config.toml")
	defer func() { ph(recover(), cfg.ErrLog) }()
	jsonbytes := must(ioutil.ReadFile("./files/sample.json")).([]byte)
	id, root, autoID, addRoot := JSONGetObjID(string(jsonbytes), "id", "DefaultRoot", PATH_DEL)
	fPln(id, root, autoID, addRoot)
}

func TestJSONScanObjects(t *testing.T) {
	cfg := c.FromFile("../build/config.toml")
	defer func() { ph(recover(), cfg.ErrLog) }()

	data := string(must(ioutil.ReadFile("./files/content.json")).([]byte))
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

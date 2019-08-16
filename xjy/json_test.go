package xjy

import (
	"io/ioutil"
	"testing"

	c "github.com/nsip/n3-client/config"
	g "github.com/nsip/n3-client/global"
)

func TestJSONObjChildren(t *testing.T) {
	cfg := c.FromFile("../build/config.toml")
	defer func() { ph(recover(), cfg.ErrLog) }()
	jsonbytes := must(ioutil.ReadFile("./files/sample.json")).([]byte)
	fPln(JSONObjChildren(string(jsonbytes)))
}

func TestJSONGetObjID(t *testing.T) {
	cfg := c.FromFile("../build/config.toml")
	defer func() { ph(recover(), cfg.ErrLog) }()
	jsonbytes := must(ioutil.ReadFile("./files/sample.json")).([]byte)
	idTag, id, root, autoID, addRoot, jsonObj := JSONObjInfo(string(jsonbytes), "DefaultRoot", g.DELIPath)
	fPln("ID Tag:    ", idTag)
	fPf("%s:         %s\n", idTag, id)
	fPln("root:      ", root)
	fPln("autoID:    ", autoID)
	fPln("added Root:", addRoot)
	fPln(jsonObj)
}

func TestJSONScanObjects(t *testing.T) {
	cfg := c.FromFile("../build/config.toml")
	defer func() { ph(recover(), cfg.ErrLog) }()

	data := string(must(ioutil.ReadFile("./files/xapi.json")).([]byte))
	mStructRecord := map[string][]string{}
	procIdx := 1
	JSONObjScan(data, "ROOT",
		func(p, id string, v []string, lastObjTuple bool) error {
			if _, ok := mStructRecord[p]; !ok {
				mStructRecord[p] = v
				fPf("S%3d ---> %-70s:: %s\n", procIdx, p, sJ(v, g.DELIChild))
				procIdx++
			}
			return nil
		},
		func(p, id string, n int, lastObjTuple bool) error {
			fPf("A%3d ---> %-70s[]%s -- [%d]\n", procIdx, p, id, n)
			procIdx++
			return nil
		},
	)
}

func TestJSONArrDiv(t *testing.T) {
	data := string(must(ioutil.ReadFile("./files/xapi10.json")).([]byte))
	arrs, rem := JSONArrDiv(string(data), 3)
	fPln(arrs[0], rem)
}

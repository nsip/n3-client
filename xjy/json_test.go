package xjy

import (
	"io/ioutil"
	"testing"

	c "../config"
	u "github.com/cdutwhu/go-util"
)

func TestJSONGetObjID(t *testing.T) {
	cfg := c.GetConfig("./config.toml", "../config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog) }()
	jsonbytes := Must(ioutil.ReadFile("./files/xapi.json")).([]byte)
	id := JSONGetObjID(string(jsonbytes), "id", "DefaultRoot", pathDel)
	fPln(id)
}

func TestJSONScanObjects(t *testing.T) {
	cfg := c.GetConfig("./config.toml", "../config/config.toml")
	defer func() { PH(recover(), cfg.Global.ErrLog) }()

	xapibytes := Must(ioutil.ReadFile("./files/xapi.json")).([]byte)
	xapi := u.Str(xapibytes)
	xapi.SetEnC()
	// fPln(s.IsJSONRootArray())

	mapStructRecord := map[string]string{}	
	if ok, jsonType, n := xapi.IsJSONRootArray(); ok {
		if jsonType == "Object" {
			prevEnd := 0
			for i := 1; i <= n; i++ {
				nextStart := u.TerOp(i == 1, 0, prevEnd+1).(int)
				json, _, end := xapi.S(nextStart, u.ALL).BracketsPos(u.BCurly, 1, 1)
				prevEnd = end + nextStart
				// if i == n - 3 {
				// 	fPln(json)
				// }

				procIdx := 1
				JSONModelInfo(json.V(), "id", "XAPI", pathDel, childDel,
					func(p, v string) {
						if prevV, ok := mapStructRecord[p]; !ok || (ok && v != prevV && u.Str(v).FieldsSeqContain(prevV, childDel)) {
							mapStructRecord[p] = v
							fPf("S%2d ---> %-80s:: %s\n", procIdx, p, v)
							procIdx++
						}						
					},
					func(p, v string, n int) {
						fPf("A ---> %5d:  %-80s[] %s -- [%d]\n", procIdx, p, v, n)
						procIdx++
					},
				)				
			}
		}
	}

	// idx := 1
	// JSONModelInfo(string(jsonbytes), "id", "DefaultRoot", pathDel, childDel,
	// 	func(p, v string) {
	// 		fPf("S ---> %5d:  %-80s:: %s\n", idx, p, v)
	// 		idx++
	// 	},
	// 	func(p, v string, n int) {
	// 		fPf("A ---> %5d:  %-80s[] %s -- [%d]\n", idx, p, v, n)
	// 		idx++
	// 	},
	// )
}

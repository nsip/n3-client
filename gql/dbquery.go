package gql

import (
	c "../config"
	g "../global"
	q "../query"
	u "github.com/cdutwhu/go-util"
	"github.com/nsip/n3-messages/n3grpc"
)

// Init :
func Init(config *c.Config) {
	PC(config == nil, fEf("Init Config"))
	Cfg = config
	g.N3clt = u.TerOp(g.N3clt == nil, n3grpc.NewClient(Cfg.RPC.Server, Cfg.RPC.Port), g.N3clt).(*n3grpc.Client)
}

// filling root, mapStruct, mapValue, mapArray
func queryObject(id string) {
	defer func() { PH(recover(), Cfg.Global.ErrLog) }()
	Init(c.GetConfig("./config.toml", "../config/config.toml"))
	root, mapStruct, mapValue, mapArray = "", map[string]string{}, map[string]string{}, map[string]int{}

	_, _, o, _ := q.Sif(id, "") //               *** get Object's Root ***
	if len(o) > 0 {
		root = o[0]
	} else {
		return
	}

	s, _, o, _ := q.Sif(root, "::") //           *** get Object's Struct ***
	for i := range s {
		mapStruct[s[i]] = o[i]
	}

	s, p, o, _ := q.Sif(id, root) //             *** get Object's Values ***
	for i := range s {
		mapValue[p[i]] = o[i]
	}

	s, p, o, _ = q.Sif(id, "ARR") //             *** get Object's array info ***
	for i := range s {
		mapArray[p[i]] = u.Str(o[i]).ToInt()
	}

	return
}

func isEndValue(path string) bool {
	_, ok := mapValue[path]
	return ok
}

func isArrStruct(path string) (string, bool) {
	if v, ok := mapStruct[path]; ok {
		if sHP(v, "[]") {
			return v[2:], true
		}
	}
	return "", false
}

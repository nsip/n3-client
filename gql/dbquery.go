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

	_, _, o, _ := q.Sif(id, "") //               *** Object's Root ***
	if len(o) > 0 {
		root = o[0]
	} else {
		return
	}

	s, _, o, _ := q.Sif(root, "::") //           *** Object's Struct ***
	for i := range s {
		mapStruct[s[i]] = o[i]
	}

	s, p, o, v := q.Sif(id, root) //             *** Object's Values ***
	for i := range s {
		mapValue[p[i]] = append(mapValue[p[i]], &valver{value: o[i], ver: v[i]})
	}

	s, p, o, _ = q.Sif(id, "ARR") //             *** Object's array count ***
	for i := range s {
		mapArray[p[i]] = u.Str(o[i]).ToInt()
	}

	return
}

func isEndValue(path string) ([]*valver, bool) {
	v, ok := mapValue[path]
	return v, ok
}

func cntChildren(path string) (cnt int) {
	if v, ok := mapStruct[path]; ok {
		cnt = len(sSpl(v, " + "))
	}
	return 
}

func isArr(path string) (string, int, bool) {
	if v, ok := mapStruct[path]; ok {
		if sHP(v, "[]") {
			return v[2:], mapArray[path], true
		}
	}
	return "", 0, false
}

func isParentArr(path string) (int, bool) {
	if p := sLI(path, "."); p > 0 {
		ppath := path[:p]
		_, cnt, ok := isArr(ppath)
		return cnt, ok
	}
	return 0, false
}

func isArrPath(path string) (int, bool) {
	xpath, arrCnt, OK := "", 1, false
	for _, seg := range sSpl(path, ".") {
		xpath += (seg + ".")
		if _, cnt, ok := isArr(xpath[:len(xpath)-1]); ok {
			arrCnt *= cnt
			OK = ok
		}
	}
	return arrCnt, OK
}

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
func queryObject(id, from string) {
	defer func() { PH(recover(), Cfg.Global.ErrLog) }()
	Init(c.GetConfig("./config.toml", "../config/config.toml"))

	fn := u.CaseAssign(from, "sif", "xapi", q.Sif, q.Xapi).(func(sp ...string) (s, p, o []string, v []int64))

	_, _, o, _ := fn(id, "") //               *** Object's Root ***
	if len(o) > 0 {
		root = o[0]
	} else {
		return
	}

	fPln("<id> :", id, "<root> :", root)

	s, _, o, _ := fn(root, "::") //           *** Object's Struct ***
	for i := range s {
		mapStruct[s[i]] = o[i]
		fPln("S -------> ", s[i], "   ", o[i])
	}

	s, p, o, _ := fn(id, "[]") //              *** Object's array count ***
	for i := range s {
		mapArray[p[i]] = u.Str(o[i]).ToInt()
		fPln("A -------> ", p[i], "   ", o[i])
	}

	s, p, o, v := fn(id, root) //             *** Object's Values ***
	for i := range s {
		mapValue[p[i]] = append(mapValue[p[i]], &valver{value: o[i], ver: v[i]})
		fPln("V -------> ", p[i], "   ", o[i], "   ", v[i])
	}

	return
}

func isLeafValue(path string) (bool, []*valver) {
	path = sRepAll(path, "[]", "")
	v, ok := mapValue[path]
	return ok, v
}

func isObject(path string) bool {
	path = sRepAll(path, "[]", "")
	_, ok1 := mapStruct[path]
	_, ok2 := mapArray[path]
	return ok1 && !ok2
}

func isArray(path string) (fArr bool, nArr int, plain bool) {
	path = sRepAll(path, "[]", "")
	nArr, fArr = mapArray[path]
	_, ok := mapStruct[path]
	return fArr, nArr, !ok
}

// func cntChildren(path, childDel string) (cnt int) {
// 	if v, ok := mapStruct[path]; ok {
// 		cnt = len(sSpl(v, childDel))
// 	}
// 	return
// }

// func containsArr(path, pathDel, childDel string) (arrNames []string, arrCnts []int, fContain bool) {
// 	if v, ok := mapStruct[path]; ok {
// 		for _, subS := range sSpl(v, childDel) {
// 			if sHP(subS, "[]") {
// 				arrNames = append(arrNames, subS[2:])
// 				arrCnts = append(arrCnts, mapArray[path+pathDel+subS[2:]])
// 				fContain = true
// 			} else {
// 				arrNames = append(arrNames, subS)
// 				arrCnts = append(arrCnts, 1)
// 			}
// 		}
// 		return arrNames, arrCnts, fContain
// 	}
// 	return nil, nil, false
// }

// func isArrPath(path, pathDel, childDel string) (int, bool) {
// 	xpath, arrCnt, OK := "", 1, false
// 	for _, seg := range sSpl(path, pathDel) {
// 		xpath += (seg + pathDel)
// 		if _, cnt, ok := containsArr(xpath[:len(xpath)-len(pathDel)], pathDel, childDel); ok {
// 			arrCnt *= cnt[0]
// 			OK = ok
// 		}
// 	}
// 	return arrCnt, OK
// }

// func isParentArr(path, del string) ([]int, bool) {
// 	if p := sLI(path, del); p > 0 {
// 		ppath := path[:p]
// 		_, cnt, ok := containsArr(ppath, del)
// 		return cnt, ok
// 	}
// 	return nil, false
// }

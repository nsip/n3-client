package gql

import (
	c "../config"
	g "../global"
	q "../query"
	"github.com/nsip/n3-messages/n3grpc"
)

// Init :
func Init(config *c.Config) {
	PC(config == nil, fEf("Init Config"))
	CFG = config
	g.N3clt = IF(g.N3clt == nil, n3grpc.NewClient(CFG.RPC.Server, CFG.RPC.Port), g.N3clt).(*n3grpc.Client)
}

func clrQueryCache() {
	mStruct = map[string]string{}
	mValue = map[string][]*valver{}
	mArray = map[string]int{}
	mIndicesList = map[string][][]int{}
	mIPathObj = map[string]string{}
	mIPathSubIPaths = map[string][]string{}	
}

// filling root, mStruct, mValue, mArray
func queryObject(id string) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	Init(c.FromFile("./config.toml", "../config/config.toml"))

	_, _, o, _ := q.Data(id, "") //              *** Object's Root ***
	if len(o) > 0 {
		root = o[0]
	} else {
		return
	}

	fPln(" ---------------------------------------------- ")

	s, p, o, v := q.Data(id, "::") //            *** Struct ***
	for i := range s {
		mStruct[p[i]] = o[i]
		// fPf("S mStruct --> %-70s%-70s%10d\n", p[i], o[i], v[i])
	}

	fPln(" ---------------------------------------------- ")

	s, p, o, v = q.Data(id, "[]") //             *** Array ***
	for i := range s {
		mArray[p[i]] = Str(o[i]).ToInt()
		// fPf("A mArray --> %-70s%-70s%10d\n", p[i], o[i], v[i])
	}
	mIndicesList = mkIndicesList()

	fPln(" ---------------------------------------------- ")

	s, p, o, v = q.Data(id, root) //             *** Values ***
	for i := range s {
		mValue[p[i]] = append(mValue[p[i]], &valver{value: o[i], ver: v[i]})
		// fPf("V mValue --> %-70s%-70s%10d\n", p[i], o[i], v[i])
	}

	return
}

func isLeafValue(path string) (bool, []*valver) {
	// path = sRepAll(path, "[]", "")
	v, ok := mValue[path]
	return ok, v
}

func isObject(path string) bool {
	path = sRepAll(path, "[]", "")
	_, ok1 := mStruct[path]
	_, ok2 := mArray[path]
	return ok1 && !ok2
}

func isArray(ipath string) (fArr bool, nArr int, plain bool) {
	nArr, fArr = mArray[ipath]
	s1, _ := Str(ipath).SplitEx(PATH_DEL, "#", "string", "int")
	path := sJ(s1.([]string), PATH_DEL)
	_, ok := mStruct[path]
	return fArr, nArr, !ok
}

func isParentArray(ipath string) (fArr bool, nArr int, plain bool) {
	ipathParent := Str(ipath).RmTailFromLast(PATH_DEL).V()
	ipathParent = Str(ipathParent).RmTailFromLast("#").V()
	return isArray(ipathParent)
}

// ipaths is sorted
func getVIdxForIPath(ipaths []string, ipath string) (idx int) {
	for _, ip := range ipaths {
		if ip == ipath {
			return
		}
		if ok, n, plain := isArray(ip); ok && plain {
			idx += n
		}
	}
	return
}

// mkIndicesList : Need <mArray>
func mkIndicesList() (rst map[string][][]int) {
	rst = make(map[string][][]int)
	if len(mArray) == 0 {
		return
	}
	keys := GetMapKeys(mArray).([]string)
	for _, k := range keys {
		s1, i2 := Str(k).SplitEx(PATH_DEL, "#", "string", "int")
		rKey := sJ(s1.([]string), PATH_DEL)
		rst[rKey] = append(rst[rKey], i2.([]int))
	}
	for k := range rst {
		SortIntArr2D(rst[k], "ASC")
	}
	return
}

// IPathListBymArr : Need <mArray>
func IPathListBymArr(path string) (rst []string) {
	if len(mArray) == 0 {
		return
	}
	for k, v := range mIndicesList {
		if k == path {
			for _, ind := range v {
				// fPln(ind)
				sind := []string{}
				for _, idx := range ind {
					sind = append(sind, fSf("%d", idx))
				}
				rst = append(rst, IArrStrJoinEx(Strs(sSpl(path, PATH_DEL)), Strs(sind), "#", PATH_DEL))
			}
		}
	}
	return
}

// SubIPathListByPath : Need <mIPathSubIPaths>
func SubIPathListByPath(path string) (rst []string) {
	if len(mIPathSubIPaths) == 0 {
		return nil
	}
	mapKeys := GetMapKeys(mIPathSubIPaths).([]string)
	for _, k := range mapKeys {
		r1, _ := Str(k).SplitEx(PATH_DEL, "#", "string", "int")
		if sJ(r1.([]string), PATH_DEL) == path {
			ipath := k
			if subIPaths, ok := mIPathSubIPaths[ipath]; ok {
				for _, subIPath := range subIPaths {
					rst = append(rst, subIPath)
				}
			}
		}
	}

	S1 := []string{}
	a2d := [][]int{}
	for _, ip := range rst {
		s1, i2 := Str(ip).SplitEx(PATH_DEL, "#", "string", "int")
		a2d = append(a2d, i2.([]int))
		S1 = s1.([]string)
	}
	SortIntArr2D(a2d, "ASC")

	rst = []string{}
	for _, arr := range a2d {
		sa := make([]string, len(arr))
		for i, a := range arr {
			sa[i] = fSf("%d", a)
		}
		rst = append(rst, IArrStrJoinEx(Strs(S1), Strs(sa), "#", PATH_DEL))
	}

	return
}

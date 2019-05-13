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

func clrQueryBuf() {
	mStruct = map[string]string{}
	mValue = map[string][]*valver{}
	mArray = map[string]int{}
	mIPathObj = map[string]string{}
	mIPathSubIPaths = map[string][]string{}	
}

// filling root, mStruct, mValue, mArray
func queryObject(id string) {
	defer func() { PH(recover(), CFG.Global.ErrLog) }()
	Init(c.FromFile("./config.toml", "../config/config.toml"))

	_, _, o, _ := q.Data(id, "") //               *** Object's Root ***
	if len(o) > 0 {
		root = o[0]
	} else {
		return
	}

	// fPln("<id> :", id, "<root> :", root)

	fPln(" ---------------------------------------------- ")

	s, _, o, _ := q.Data(root, "::") //            *** Struct ***
	for i := range s {
		if stru, ok := mStruct[s[i]]; !ok || (ok && stru != o[i] && Str(o[i]).FieldsSeqCtn(stru, CHILD_DEL)) {
			mStruct[s[i]] = o[i]
			fPf("S mStruct---> %-70s : %s\n", s[i], o[i])
		}
		// mStruct[s[i]] = o[i]
		// fPln("S mStruct -------> ", s[i], "   ", o[i])
	}

	fPln(" ---------------------------------------------- ")

	s, p, o, _ := q.Data(id, "[]") //              *** Array count ***
	for i := range s {
		mArray[p[i]] = Str(o[i]).ToInt()
		fPf("A mArray ---> %-70s : %s\n", p[i], o[i])
	}

	fPln(" ---------------------------------------------- ")

	s, p, o, v := q.Data(id, root) //              *** Values ***
	for i := range s {
		mValue[p[i]] = append(mValue[p[i]], &valver{value: o[i], ver: v[i]})
		fPf("V mValue ---> %-70s%-90s%10d\n", p[i], o[i], v[i])
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

// // PermuIndices :
// func PermuIndices(maxIndices []int) (rst [][]int) {

// 	OriL := len(maxIndices)
// 	MaxIndices := []int{1, 1, 1, 1, 1, 1, 1, 1, 1} //                              *** MAX is 9
// 	PC(OriL > len(MaxIndices), fEf("Only MAX LENGTH <%d> maxIndices is supported", len(MaxIndices)))
// 	copy(MaxIndices, maxIndices)

// 	C := 1
// 	for _, mIdx := range MaxIndices {
// 		C *= mIdx
// 	}
// 	rst = make([][]int, C)

// 	I, MI := 0, MaxIndices
// 	for i0 := 1; i0 <= MI[0]; i0++ {
// 		for i1 := 1; i1 <= MI[1]; i1++ {
// 			for i2 := 1; i2 <= MI[2]; i2++ {
// 				for i3 := 1; i3 <= MI[3]; i3++ {
// 					for i4 := 1; i4 <= MI[4]; i4++ {
// 						for i5 := 1; i5 <= MI[5]; i5++ {
// 							for i6 := 1; i6 <= MI[6]; i6++ {
// 								for i7 := 1; i7 <= MI[7]; i7++ {
// 									for i8 := 1; i8 <= MI[8]; i8++ {
// 										rst[I] = []int{i0, i1, i2, i3, i4, i5, i6, i7, i8}
// 										I++
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}

// 	for i := 0; i < C; i++ {
// 		rst[i] = rst[i][:OriL]
// 	}

// 	return
// }

// mArrayIndicesList : Need <mArray>
func mArrayIndicesList() (rst map[string][][]int) {
	rst = make(map[string][][]int)
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
	for k, v := range mArrayIndicesList() {
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

// ------------------------------------------------------------------------------------------------------------

// func cntChildren(path, childDel string) (cnt int) {
// 	if v, ok := mStruct[path]; ok {
// 		cnt = len(sSpl(v, childDel))
// 	}
// 	return
// }

// func containsArr(path, pathDel, childDel string) (arrNames []string, arrCnts []int, fContain bool) {
// 	if v, ok := mStruct[path]; ok {
// 		for _, subS := range sSpl(v, childDel) {
// 			if sHP(subS, "[]") {
// 				arrNames = append(arrNames, subS[2:])
// 				arrCnts = append(arrCnts, mArray[path+pathDel+subS[2:]])
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

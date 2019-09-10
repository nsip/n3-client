package gql

import (
	g "github.com/nsip/n3-client/global"
	q "github.com/nsip/n3-client/query"
)

func clrQueryCache() {
	mStruct = map[string]string{}
	mValue = map[string][]*valver{}
	mArray = map[string]int{}
	mIndicesList = map[string][][]int{}
	mIPathObj = map[string]string{}
	mIPathSubIPaths = map[string][]string{}
}

// filling root, mStruct, mValue, mArray
func queryObject(ctx, id string) {
	defer func() { ph(recover(), g.Cfg.ErrLog) }()

	_, _, o, _ := q.Data(ctx, id, "") //         *** Object's Root ***
	if o != nil && len(o) > 0 {
		root = o[0]
		// fPf("Root --> %s\n", o[0])
	} else {
		return
	}

	// fPln(" ---------------------------------------------- ")

	s, p, o, v := q.Data(ctx, id, "::") //       *** Struct ***
	if s != nil {
		for i := range s {
			mStruct[p[i]] = o[i]
			// fPf("S mStruct --> %-70s%-70s%10d\n", p[i], o[i], v[i])
		}
	}

	// fPln(" ---------------------------------------------- ")

	s, p, o, v = q.Data(ctx, id, "[]") //        *** Array ***
	if s != nil {
		for i := range s {
			mArray[p[i]] = S(o[i]).ToInt()
			// fPf("A mArray --> %-70s%-70s%10d\n", p[i], o[i], v[i])
		}
		mIndicesList = mkIndicesList()
	}

	// fPln(" ---------------------------------------------- ")

	s, p, o, v = q.Data(ctx, id, root) //        *** Values ***
	if s != nil {
		for i := range s {
			mValue[p[i]] = append(mValue[p[i]], &valver{value: o[i], ver: v[i]})
			// fPf("V mValue --> %-70s%-70s%10d\n", p[i], o[i], v[i])
		}
	}

	return
}

func isLeafValue(path string) (bool, []*valver) {
	// path = sRepAll(path, "[]", "")
	v, ok := mValue[path]
	return ok, v
}

func isObject(path string) bool {
	path = S(path).Replace("[]", "").V()
	_, ok1 := mStruct[path]
	_, ok2 := mArray[path]
	return ok1 && !ok2
}

func isArray(ipath string) (fArr bool, nArr int, plain bool) {
	nArr, fArr = mArray[ipath]
	s1, _ := S(ipath).SplitEx(g.DELIPath, "#", "string", "int")
	path := sJ(s1.([]string), g.DELIPath)
	_, ok := mStruct[path]
	return fArr, nArr, !ok
}

func isParentArray(ipath string) (fArr bool, nArr int, plain bool) {
	ipathParent := S(ipath).RmTailFromLast(g.DELIPath).V()
	ipathParent = S(ipathParent).RmTailFromLast("#").V()
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
		s1, i2 := S(k).SplitEx(g.DELIPath, "#", "string", "int")
		rKey := sJ(s1.([]string), g.DELIPath)
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
				rst = append(rst, IArrStrJoinEx(Ss(sSpl(path, g.DELIPath)), Ss(sind), "#", g.DELIPath))
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
		r1, _ := S(k).SplitEx(g.DELIPath, "#", "string", "int")
		if sJ(r1.([]string), g.DELIPath) == path {
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
		s1, i2 := S(ip).SplitEx(g.DELIPath, "#", "string", "int")
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
		rst = append(rst, IArrStrJoinEx(Ss(S1), Ss(sa), "#", g.DELIPath))
	}

	return
}

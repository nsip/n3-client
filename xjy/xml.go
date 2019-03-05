package xjy

import (
	u "github.com/cdutwhu/go-util"
)

// XMLScanObjects is ( ONLY LIKE  <SchoolInfo RefId="D3F5B90C-D85D-4728-8C6F-0D606070606C"> )
func XMLScanObjects(xml, idmark string) (ids, objtags []string, posarr []int) {
	idmark = u.Str(idmark).MkPrefix(" ")
	idmark = u.Str(idmark).MkSuffix("=")
	lengthID := len(idmark)
	pLastAbs := 0
LOOKFOROBJ:
	if p := sI(xml[pLastAbs:], idmark); p > 0 {
		if op := sLI(xml[pLastAbs:pLastAbs+p], "<"); op >= 0 {
			obj := xml[pLastAbs : pLastAbs+p][op+1:]
			objtags = append(objtags, obj)
			ps := pLastAbs + op
			posarr = append(posarr, ps)
		}
		if ip := sI(xml[pLastAbs+p:], ">"); ip > 0 {
			id := xml[pLastAbs+p+lengthID : pLastAbs+p+ip]
			id = sT(id, "\"")
			ids = append(ids, id)
		}
		pLastAbs += (p + lengthID)
		goto LOOKFOROBJ
	}
	return
}

// XMLObjStrByID is
func XMLObjStrByID(xml, idmark, rid string) string {
	ids, objtags, posarr := XMLScanObjects(xml, idmark)
	for i, id := range ids {
		if id == rid {
			if i != len(ids)-1 {
				return sTR(xml[posarr[i]:posarr[i+1]], " \t\r\n")
			}
			/* last object */
			endtag := "</" + objtags[i] + ">"
			if end := sI(xml[posarr[i]:], endtag); end > 0 {
				return sTR(xml[posarr[i]:posarr[i]+end+len(endtag)], " \t\r\n")
			}
		}
	}
	return ""
}

// XMLEleStrByTag is (should only be used in one object string)
// func XMLEleStrByTag(xml, tag string) string {
// 	s, s1 := sI(xml, fSpf("<%s>", tag)), sI(xml, fSpf("<%s ", tag))
// 	if s1 > s {
// 		s = s1
// 	}
// 	if s >= 0 {
// 		if e := sI(xml[s:], fSpf("</%s>", tag)); e > 0 {
// 			return xml[s : s+e+len(tag)+3]
// 		}
// 		PE(fEf("Not a valid XML"))
// 	}
// 	return ""
// }

// XMLEleStrByTag : (index from 1)
func XMLEleStrByTag(xml, tag string, idx int) string {
	startNext, cnt := 0, 0
AGAIN:
	xml = xml[startNext:]
	s, s1 := sI(xml, fSf("<%s>", tag)), sI(xml, fSf("<%s ", tag))
	if s1 > s {
		s = s1
	}
	if s >= 0 {
		if peR := sI(xml[s:], fSf("</%s>", tag)); peR > 0 {
			startNext = s + peR + len(tag) + 3
			cnt++
			if idx == cnt {
				return xml[s:startNext]
			}
			goto AGAIN
		}
		PE(fEf("Invalid XML"))
	}
	return ""
}

// XMLFindAttributes is (ONLY LIKE  <SchoolInfo RefId="D3F5B90C-D85D-4728-8C6F-0D606070606C" Type="LGL">)
func XMLFindAttributes(xmlele, del string) (attributes, attriValues []string, attributeList string) { /* 'map' may cause mis-order, so use slice */
	l := len(xmlele)
	if l == 0 || xmlele[0] != '<' || xmlele[l-1] != '>' {
		PE(fEf("Not a valid XML section"))
		return nil, nil, ""
	}

	tag := xmlele[sLI(xmlele, "</")+2 : l-1]
	if eol := sI(xmlele, "\">") + 1; xmlele[len(tag)+1] == ' ' && eol > len(tag) { /* has attributes */
		kvs := sFF(xmlele[len(tag)+2:eol], func(c rune) bool { return c == ' ' })
		for _, kv := range kvs {
			kvstrs := sFF(kv, func(c rune) bool { return c == '=' })
			attributes = append(attributes, ("-" + kvstrs[0])) /* mark '-' before attribute for differentiating child */
			attriValues = append(attriValues, u.Str(kvstrs[1]).RmQuotes())
		}
	}
	return attributes, attriValues, sJ(attributes, del)
}

// XMLFindChildren : (NOT search grandchildren)
func XMLFindChildren(xmlele, id, del string) (uids, children []string, childList string, arrCnt int) {
	l := len(xmlele)
	if l == 0 || xmlele[0] != '<' || xmlele[l-1] != '>' {
		fPln(xmlele)
		PE(fEf("Not a valid XML section"))
		return nil, nil, "nil", -1
	}

	skip, childpos, level, inflag := false, []int{}, 0, false
	for i, c := range xmlele[1:] { // skip the first '<'
		i++

		if c == '<' && xmlele[i:i+4] == "<!--" {
			skip = true
		}
		if c == '>' && xmlele[i-2:i+1] == "-->" {
			skip = false
		}
		if skip {
			continue
		}

		if c == '<' && xmlele[i+1] != '/' {
			level++
		}
		if c == '<' && xmlele[i+1] == '/' {
			level--
			if level == 0 {
				inflag = false
			}
		}

		if level == 1 {
			if !inflag {
				childpos = append(childpos, i+1)
				inflag = true
			}
		}
	}

	for _, p := range childpos {
		pe, peA := sI(xmlele[p:], ">"), sI(xmlele[p:], " ")
		if peA > 0 && peA < pe {
			pe = peA
		}
		child := xmlele[p : p+pe]
		children = append(children, child)
		uids = append(uids, id)
	}

	if len(children) > 1 && u.Strs(children).ToG().AllAreIdentical() {
		return uids, children, fSf("[]%s", children[0]), len(children) /* get array count from db, not here. */
	}

	return uids, children, sJ(children, del), 0
}

// XMLYieldArrInfo :
func XMLYieldArrInfo(xmlstr string, ids, objs []string, mapkeyprefix, pathDel, childDel string, eleObjIDArrcnts *[]pathIDn) {
	if len(mapkeyprefix) > 0 {
		mapkeyprefix += pathDel
	}
	for i, obj := range objs {
		curPath := mapkeyprefix + obj

		xmlele := XMLEleStrByTag(xmlstr, obj, 1)
		uids, children, _, arrCnt := XMLFindChildren(xmlele, ids[i], childDel) /* uniform ids, children */
		attributes, _, _ := XMLFindAttributes(xmlele, childDel)                /* attributes */

		/* array children info */
		if arrCnt > 0 {
			(*eleObjIDArrcnts) = append((*eleObjIDArrcnts), pathIDn{arrPath: curPath + pathDel + children[0], objID: ids[i], arrCnt: arrCnt})
		}

		if len(children) == 0 && len(attributes) == 0 { /* attributes */
			continue
		} else {
			XMLYieldArrInfo(xmlele, uids, children, curPath, pathDel, childDel, eleObjIDArrcnts) /* recursive */
		}
	}
}

// XMLFamilyTree is (We pack attributes in return map, value like '-...')
func XMLFamilyTree(xmlstr string, ids, objs []string, skipNoChild bool, mapkeyprefix, pathDel, childDel string, mapEleChildlist *map[string]string) {
	if len(mapkeyprefix) > 0 {
		mapkeyprefix += pathDel
	}
	for i, obj := range objs {
		curPath := mapkeyprefix + obj

		if _, ok := (*mapEleChildlist)[curPath]; ok {
			continue /* ONLY keep one identical path's children */
		}

		xmlele := XMLEleStrByTag(xmlstr, obj, 1)
		uids, children, childlist, _ := XMLFindChildren(xmlele, ids[i], childDel) /* uniform ids, children */
		attributes, _, attributeList := XMLFindAttributes(xmlele, childDel)       /* attributes */

		/* attributes */
		if len(attributes) > 0 {
			(*mapEleChildlist)[curPath] = attributeList + childDel
			if len(children) == 0 {
				(*mapEleChildlist)[curPath] += "#content"
			}
		}

		/* children */
		if skipNoChild {
			if len(children) > 0 {
				(*mapEleChildlist)[curPath] += childlist
			}
		} else {
			(*mapEleChildlist)[curPath] += childlist
		}

		if len(children) == 0 && len(attributeList) == 0 { /* attributes */
			continue
		} else {
			XMLFamilyTree(xmlele, uids, children, skipNoChild, curPath, pathDel, childDel, mapEleChildlist) /* recursive */
		}
	}
}

// pathIDn : array's path, object ID, array's count
type pathIDn struct {
	arrPath string
	objID   string
	arrCnt  int
}

// XMLModelInfo :
func XMLModelInfo(xmlstr, ObjIDMark, pathDel, childDel string, OnStruFetch func(string, string), OnArrFetch func(string, string, int)) {
	ids, objs, _ := XMLScanObjects(xmlstr, ObjIDMark)

	mapEleChildlist := &map[string]string{}
	XMLFamilyTree(xmlstr, ids, objs, true, "", pathDel, childDel, mapEleChildlist)
	for k, v := range *mapEleChildlist {
		OnStruFetch(k, v)
	}

	eleObjIDArrcnts := &[]pathIDn{}
	XMLYieldArrInfo(xmlstr, ids, objs, "", pathDel, childDel, eleObjIDArrcnts)
	for _, c := range *eleObjIDArrcnts {
		OnArrFetch(c.arrPath, c.objID, c.arrCnt)
	}
}

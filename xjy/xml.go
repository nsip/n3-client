package xjy

// XMLScanObjects is ( ONLY LIKE  <SchoolInfo RefId="D3F5B90C-D85D-4728-8C6F-0D606070606C"> )
func XMLScanObjects(xml, idmark string) (ids, objtags []string, starts, ends []int) {
	L := Str(xml).L()
	nXML, pNextStart, pLastPos := XMLSegsCount(xml), 0, 0
	for i := 1; i <= nXML; i++ {

		// fPln(i)
		// if i == 2 {
		// 	fPln("debug break", L)
		// }

		tag, thisXML, leftR, rightR := XMLSegPos(Str(xml).S(pNextStart, ALL).V(), 1, 1)
		objtags = append(objtags, tag)

		attri, attrivalues := XMLAttributes(thisXML, "")
		for j := 0; j < len(attri); j++ {
			if attri[j] == idmark {
				ids = append(ids, attrivalues[j])
				break
			}
		}

		pLastPos = pNextStart + leftR
		starts = append(starts, pLastPos)
		for j := pLastPos + rightR; j >= 0; j-- {
			j = IF(j >= L, L-1, j).(int)
			if Str(xml).C(j) == '>' {
				ends = append(ends, j)
				break
			}
		}

		pNextStart += rightR
	}
	return
}

// XMLObjStrByID is
func XMLObjStrByID(xml, idmark, rid string) string {
	ids, objtags, starts, _ := XMLScanObjects(xml, idmark)
	XML, nIDs := Str(xml), len(ids)
	for i, id := range ids {
		if id == rid {
			if i != nIDs-1 {
				return XML.S(starts[i], starts[i+1]).T(BLANK).V()
			}
			/* last object */
			endtag := "</" + objtags[i] + ">"
			if end := XML.S(starts[i], ALL).Idx(endtag); end > 0 {
				return XML.S(starts[i], starts[i]+end+Str(endtag).L()).T(BLANK).V()
			}
		}
	}
	return ""
}

// XMLInfoScan :
func XMLInfoScan(xmlstr, objIDMark, PATHDEL string,
	OnStructFetch func(string, string, []string, bool),
	OnArrFetch func(string, string, int, bool)) (IDs []string) {

	ids, objs, starts, ends := XMLScanObjects(xmlstr, objIDMark)
	nObj := len(ids)

	for i := 0; i < nObj; i++ {

		id, _, xml := ids[i], objs[i], Str(xmlstr).S(starts[i], ends[i]+1).V()
		mFT, mArr := XMLCntInfo(xml, "", PATHDEL, id, nil)

		for k, v := range *mArr { //                      *** only need >= 2 to be xml array ***
			if v.Count <= 1 {
				delete(*mArr, k)
			}
		}

		j, lFT, lArr := 0, len(*mFT), len(*mArr)
		for k, v := range *mFT {
			j++
			OnStructFetch(k, id, v, (j == lFT)) //        *** last tuple flag ***
		}

		j = 0
		for k, v := range *mArr {
			j++
			OnArrFetch(k, v.ID, v.Count, (j == lArr)) //  *** last tuple flag ***
		}
	}
	return ids
}

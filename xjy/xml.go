package xjy

import "github.com/google/uuid"

// XMLScanObjects is ( ONLY LIKE  <SchoolInfo RefId="D3F5B90C-D85D-4728-8C6F-0D606070606C"> )
func XMLScanObjects(xml string) (objtags, ids, idtags []string, starts, ends []int) {
	L := S(xml).L()
	nXML, pNextStart, pLastPos := XMLSegsCount(xml), 0, 0
	for i := 1; i <= nXML; i++ {

		tag, thisXML, leftR, rightR := XMLSegPos(S(xml).S(pNextStart, ALL).V(), 1, 1)
		objtags = append(objtags, tag)
		attri, attrivalues := XMLAttributes(thisXML)

		// DONE:
		mMarkUUID := make(map[string]string)
		sidtag := "I will be the shortest length ID mark, the shortest ID Mark is what we wanted"
		for j := 0; j < len(attri); j++ {
			Attri, AttriV := S(attri[j]), S(attrivalues[j])
			if IArrEleIn(Attri, Ss{"id", "ID", "Id"}) || (Attri.ToLower() == S(tag).ToLower()+"id") {
				if AttriV.IsUUID() {
					sidtag = Attri.V()
					mMarkUUID[Attri.V()] = AttriV.V()
					break
				}
			}
			if (Attri.HS("ID") || Attri.HS("id") || Attri.HS("Id")) && AttriV.IsUUID() {
				idtag := Attri.V()
				mMarkUUID[idtag] = AttriV.V()
				sidtag = IF(len(idtag) < len(sidtag), idtag, sidtag).(string)
			}
		}
		if id, ok := mMarkUUID[sidtag]; ok && id != "" {
			idtags, ids = append(idtags, sidtag), append(ids, id)
		} else {
			idtags, ids = append(idtags, "AutoID"), append(ids, uuid.New().String())
		}
		//

		pLastPos = pNextStart + leftR
		starts = append(starts, pLastPos)
		for j := pLastPos + rightR; j >= 0; j-- {
			j = IF(j >= L, L-1, j).(int)
			if S(xml).C(j) == '>' {
				ends = append(ends, j)
				break
			}
		}

		pNextStart += rightR
	}
	return
}

// XMLObjStrByID is
func XMLObjStrByID(xml, rid string) string {
	objtags, ids, _, starts, _ := XMLScanObjects(xml)
	XML, nIDs := S(xml), len(ids)
	for i, id := range ids {
		if id == rid {
			if i != nIDs-1 {
				return XML.S(starts[i], starts[i+1]).T(BLANK).V()
			}
			/* last object */
			endtag := "</" + objtags[i] + ">"
			if end := XML.S(starts[i], ALL).Idx(endtag); end > 0 {
				return XML.S(starts[i], starts[i]+end+S(endtag).L()).T(BLANK).V()
			}
		}
	}
	return ""
}

// XMLInfoScan :
func XMLInfoScan(xmlstr, PATHDEL string,
	OnStructFetch func(string, string, []string, bool) error,
	OnArrayFetch func(string, string, int, bool) error) ([]string, []string, error) {

	objs, ids, _, starts, ends := XMLScanObjects(xmlstr)
	nObj := len(ids)

	for i := 0; i < nObj; i++ {
		id, _, xml := ids[i], objs[i], S(xmlstr).S(starts[i], ends[i]+1).V()
		mFT, mArr := XMLCntInfo(xml, "", PATHDEL, id, nil)

		for k, v := range *mArr { //                        *** only need >= 2 to be xml array ***
			if v.Count <= 1 {
				delete(*mArr, k)
			}
		}

		j, lFT, lArr := 0, len(*mFT), len(*mArr)
		for k, v := range *mFT {
			j++
			if e := OnStructFetch(k, id, v, (j == lFT)); e != nil { //          *** last tuple flag ***
				return nil, nil, e
			}
		}

		j = 0
		for k, v := range *mArr {
			j++
			if e := OnArrayFetch(k, v.ID, v.Count, (j == lArr)); e != nil { //  *** last tuple flag ***
				return nil, nil, e
			}
		}
	}
	return ids, objs, nil
}

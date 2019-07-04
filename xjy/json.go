package xjy

import (
	g "../global"
	"github.com/google/uuid"
)

// JSONGetObjID : (must have TOP-LEVEL "id" like `"id": "6690e6c9-3ef0-4ed3-8b37-7f3964730bee",` )
func JSONGetObjID(jsonstr, idmark, dfltRoot, pDel string) (id, root string, autoID, addedRoot bool) {
	root, addedRoot, newJSON := JSONWrapRoot(jsonstr, dfltRoot)
	jsonstr = IF(addedRoot, newJSON, jsonstr).(string)
	id, _ = JSONXPathValue(jsonstr, root+pDel+idmark, pDel, []int{1, 1}...)
	if id == "" {
		autoID, id = true, uuid.New().String()
	}
	return
}

// JSONModelInfo :
func JSONModelInfo(jsonstr, ObjIDMark, dfltRoot, pDel string,
	OnStruFetch func(string, string, []string, bool),
	OnArrFetch func(string, string, int, bool)) (string, string) {

	id, root, _, addedRoot := JSONGetObjID(jsonstr, ObjIDMark, dfltRoot, pDel) //                  *** find ID Value by ObjIDMark ***
	id = S(id).RmQuotes(QDouble).V()

	mFT, mArr := JSONArrInfo(jsonstr, IF(addedRoot, dfltRoot, "").(string), pDel, id, nil)
	j, lFT, lArr := 0, len(*mFT), len(*mArr)
	for k, v := range *mFT {
		j++
		OnStruFetch(k, id, v, (j == lFT))
	}

	j = 0
	for k, v := range *mArr {
		j++
		OnArrFetch(k, v.ID, v.Count, (j == lArr))
	}
	return id, root
}

// JSONObjScan :
func JSONObjScan(json, idmark, dfltRoot string,
	OnStruFetch func(string, string, []string, bool),
	OnArrFetch func(string, string, int, bool)) (IDs, Objs []string) {

	if ok, eleType, n, eles := IsJSONArray(json); ok {
		if eleType == JT_OBJ {
			for i := 1; i <= n; i++ {
				id, root := JSONModelInfo(eles[i-1], idmark, dfltRoot, g.PATH_DEL, OnStruFetch, OnArrFetch)
				IDs = append(IDs, id)
				Objs = append(Objs, root)
			}
		}
	} else {
		id, root := JSONModelInfo(json, idmark, dfltRoot, g.PATH_DEL, OnStruFetch, OnArrFetch)
		IDs = append(IDs, id)
		Objs = append(Objs, root)
	}
	return
}

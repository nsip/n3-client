package xjy

import (
	"github.com/google/uuid"
)

// JSONGetObjID : (must have TOP-LEVEL "id" like `"id": "6690e6c9-3ef0-4ed3-8b37-7f3964730bee",` )
func JSONGetObjID(jsonstr, idmark, dfltRoot, pDel string) (id string, autoID, addRoot bool) {
	root, addRoot, newJSON := JSONWrapRoot(jsonstr, dfltRoot)
	jsonstr = IF(addRoot, newJSON, jsonstr).(string)
	id, _ = JSONXPathValue(jsonstr, root+pDel+idmark, pDel, []int{1, 1}...)
	if id == "" {
		autoID, id = true, uuid.New().String()
	}
	return
}

// JSONModelInfo :
func JSONModelInfo(jsonstr, ObjIDMark, dfltRoot, pDel string,
	OnStruFetch func(string, string, []string, bool),
	OnArrFetch func(string, string, int, bool)) (ID string) {

	id, _, addRoot := JSONGetObjID(jsonstr, ObjIDMark, dfltRoot, pDel) //                  *** find ID Value by ObjIDMark ***
	id = Str(id).RmQuotes(QDouble).V()

	mFT, mArr := JSONArrInfo(jsonstr, IF(addRoot, dfltRoot, "").(string), pDel, id, nil)
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
	return id
}

// JSONObjScan :
func JSONObjScan(json, idmark, dfltRoot string,
	OnStruFetch func(string, string, []string, bool),
	OnArrFetch func(string, string, int, bool)) (IDs []string) {

	if ok, eleType, n, eles := IsJSONArray(json); ok {
		if eleType == JT_OBJ {
			for i := 1; i <= n; i++ {
				IDs = append(IDs, JSONModelInfo(eles[i-1], idmark, dfltRoot, PATH_DEL, OnStruFetch, OnArrFetch))
			}
		}
	} else {
		IDs = append(IDs, JSONModelInfo(json, idmark, dfltRoot, PATH_DEL, OnStruFetch, OnArrFetch))
	}
	return
}

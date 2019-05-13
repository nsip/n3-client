package xjy

import (
	"github.com/google/uuid"
)

// JSONGetObjID : (must have TOP-LEVEL "id" like `"id": "6690e6c9-3ef0-4ed3-8b37-7f3964730bee",` )
func JSONGetObjID(jsonstr, idmark, defaultRoot, pDel string) (id string, autoID bool) {
	root, ext, newJSON := JSONWrapRoot(jsonstr, defaultRoot)
	jsonstr = IF(ext, newJSON, jsonstr).(string)
	id, _ = JSONXPathValue(jsonstr, root+pDel+idmark, pDel, []int{1, 1}...)
	if id == "" {
		autoID, id = true, uuid.New().String()
	}
	return
}

// JSONModelInfo :
func JSONModelInfo(jsonstr, ObjIDMark, defaultRoot, pDel string,
	OnStruFetch func(string, []string),
	OnArrFetch func(string, string, int)) (ID string) {

	id, _ := JSONGetObjID(jsonstr, ObjIDMark, defaultRoot, pDel) //                  *** find ID Value by ObjIDMark ***
	id = Str(id).RmQuotes(QDouble).V()

	mapFT, mapArrInfo := JSONArrInfo(jsonstr, defaultRoot, pDel, id, nil)
	for k, v := range *mapFT {
		OnStruFetch(k, v)
	}
	// fPln()
	for k, v := range *mapArrInfo {
		OnArrFetch(k, v.ID, v.Count)
	}
	return id
}

// JSONObjScan :
func JSONObjScan(json, idmark, defaultRoot string, OnStruFetch func(p string, v []string), OnArrFetch func(p, v string, n int)) (IDs []string) {
	if ok, eleType, n, eles := IsJSONArray(json); ok {
		if eleType == JT_OBJ {
			for i := 1; i <= n; i++ {
				id := JSONModelInfo(eles[i-1], idmark, defaultRoot, PATH_DEL, OnStruFetch, OnArrFetch)
				IDs = append(IDs, id)
			}
		}
	} else {
		id := JSONModelInfo(json, idmark, defaultRoot, PATH_DEL, OnStruFetch, OnArrFetch)
		IDs = append(IDs, id)
	}
	return
}

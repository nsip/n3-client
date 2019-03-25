package xjy

import (
	u "github.com/cdutwhu/go-util"
)

// JSONGetObjID : (must have top-level "id" like `"id": "6690e6c9-3ef0-4ed3-8b37-7f3964730bee",` )
func JSONGetObjID(jsonstr, idmark, defaultRoot, pathDel string) (id string) {
	jsonStr := u.Str(jsonstr)
	root, ext, newJSON := jsonStr.JSONRootEx(defaultRoot)
	jsonStr = u.TerOp(ext, u.Str(newJSON), jsonStr).(u.Str)
	id, _, _, _ = jsonStr.JSONXPathValue(root+pathDel+idmark, pathDel, []int{1, 1}...)
	return
}

// JSONModelInfo :
func JSONModelInfo(jsonstr, ObjIDMark, defaultRoot, pathDel, childDel string,
	OnStruFetch func(string, string), OnArrFetch func(string, string, int)) (addRoot bool) {

	jsonStr := u.Str(jsonstr)
	root, ext, newJSON := jsonStr.JSONRootEx(defaultRoot)
	jsonStr = u.TerOp(ext, u.Str(newJSON), jsonStr).(u.Str)

	id := JSONGetObjID(jsonstr, ObjIDMark, pathDel, defaultRoot) // *** find ID Value by ObjIDMark ***
	id = u.Str(id).RmQuotes().V()

	mapFT, mapArrInfo := jsonStr.JSONArrInfo(root, pathDel, id, nil)
	for k, v := range *mapFT {
		OnStruFetch(k, sJ(v, childDel))
	}
	for k, v := range *mapArrInfo {
		OnArrFetch(k, v.ID, v.Count)
	}
	return ext
}

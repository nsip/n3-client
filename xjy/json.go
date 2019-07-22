package xjy

import (
	g "../global"
	"github.com/google/uuid"
)

// JSONObjInfo : (must have TOP-LEVEL "ID" like `"ID": "6690e6c9-3ef0-4ed3-8b37-7f3964730bee",` )
func JSONObjInfo(json, dfltRoot, pDel string) (IDTag, ID, root string, autoID, addedRoot bool, jsonObj string) {
	root, addedRoot, newJSON := JSONWrapRoot(json, dfltRoot)
	jsonObj = IF(addedRoot, newJSON, json).(string)

	jsonContent, _ := JSONChildValue(jsonObj, root)
	mMarkUUID := make(map[string]string)
	sidtag := "I will be the shortest length ID mark, the shortest ID Mark is what we wanted"
	for _, child := range JSONObjChildren(jsonContent) {
		Child := S(child)
		if IArrEleIn(child, Ss([]string{"id", "ID", "Id"})) || (Child.ToLower() == S(dfltRoot).ToLower()+"id") {
			if id, _ := JSONXPathValue(jsonObj, root+pDel+child, pDel, []int{1, 1}...); S(id).IsUUID() {
				sidtag = child
				mMarkUUID[child] = id
				break
			}
		}
		if !Child.HP("[]") && (Child.HS("id") || Child.HS("ID") || Child.HS("Id")) {
			if id, _ := JSONXPathValue(jsonObj, root+pDel+child, pDel, []int{1, 1}...); S(id).IsUUID() {
				mMarkUUID[child] = id
				sidtag = IF(len(child) < len(sidtag), child, sidtag).(string)
			}
		}
	}
	if id, ok := mMarkUUID[sidtag]; ok && id != "" {
		IDTag, ID, autoID = sidtag, id, false
	} else {
		IDTag, ID, autoID = "AutoID", uuid.New().String(), true
	}

	// fPln(IDTag, ID, autoID)

	return
}

// JSONModelInfo :
func JSONModelInfo(json, dfltRoot, pDel string,
	OnStruFetch func(string, string, []string, bool),
	OnArrFetch func(string, string, int, bool)) (string, string) {

	_, id, root, _, addedRoot, _ := JSONObjInfo(json, dfltRoot, pDel) //   *** find ID Value ***
	id = S(id).RmQuotes(QDouble).V()

	mFT, mArr := JSONArrInfo(json, IF(addedRoot, dfltRoot, "").(string), pDel, id, nil)
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
func JSONObjScan(json, dfltRoot string,
	OnStruFetch func(string, string, []string, bool),
	OnArrFetch func(string, string, int, bool)) (IDs, Objs []string) {

	if ok, eleType, n, eles := IsJSONArray(json); ok {
		if eleType == JT_OBJ {
			for i := 1; i <= n; i++ {
				id, root := JSONModelInfo(eles[i-1], dfltRoot, g.DELIPath, OnStruFetch, OnArrFetch)
				IDs, Objs = append(IDs, id), append(Objs, root)
			}
		}
	} else {
		id, root := JSONModelInfo(json, dfltRoot, g.DELIPath, OnStruFetch, OnArrFetch)
		IDs, Objs = append(IDs, id), append(Objs, root)
	}
	return
}

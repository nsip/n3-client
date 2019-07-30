package xjy

import (
	"io/ioutil"

	g "github.com/nsip/n3-client/global"
	pp "github.com/nsip/n3-client/preprocess"
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
		if IArrEleIn(child, Ss{"id", "ID", "Id"}) || (Child.ToLower() == S(dfltRoot).ToLower()+"id") {
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

// JSONArrDiv :
func JSONArrDiv(json string, nDiv int) (jsonarrs []string, rem bool) {
	if ok, eleType, n, eles := IsJSONArray(json); ok {
		if eleType == JT_OBJ {
			nPer, nRem := n/nDiv, n%nDiv

			var lows, highs []int
			if nRem != 0 {
				lows, highs = make([]int, nDiv+1), make([]int, nDiv+1)
				for i := 0; i < nDiv; i++ {
					lows[i] = nPer * i
					highs[i] = nPer*(i+1) - 1
				}
				lows[nDiv] = highs[nDiv-1] + 1
				highs[nDiv] = n - 1
			} else {
				lows, highs = make([]int, nDiv), make([]int, nDiv)
				for i := 0; i < nDiv; i++ {
					lows[i] = nPer * i
					highs[i] = nPer*(i+1) - 1
				}
			}

			nPart := IF(nRem == 0, nDiv, nDiv+1).(int)
			for i := 0; i < nPart; i++ {
				l, h := lows[i], highs[i]
				jsonarr := ""
				for j := l; j <= h; j++ {
					jsonarr += eles[j] + ",\n"
				}
				jsonarr = "[" + jsonarr[:len(jsonarr)-2] + "]"
				pc(!IsJSON(jsonarr), fEf("JSONArrDiv result error"))
				jsonarr = prepJSON(jsonarr)
				jsonarrs = append(jsonarrs, jsonarr)
			}

			rem = nRem != 0
		}
	}
	return
}

func prepJSON(json string) string {

	// *** format json ***

	// json = pp.FmtJSONStr(json, "../preprocess/util/", "./")
	ioutil.WriteFile("../build/debug_pub/in.json", []byte(json), 0666)
	json = pp.FmtJSONFile("../../build/debug_pub/in.json", "../preprocess/util/", "./")
	ioutil.WriteFile("../build/debug_pub/infmt.json", []byte(json), 0666)

	// *** ': null' => ': "null"' ***
	json = S(json).Replace(`": null`, `": "null"`).V()

	// *** dealing with colon ***
	if pp.HasColonInValue(json) {
		json = pp.RplcValueColons(json)
	}

	// *** convert to ASCII ***
	if ascii, ajson := UTF8ToASCII(json); !ascii {
		fPln("is utf8")
		return ajson
	}
	return json
}

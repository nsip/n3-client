package xjy

import g "../global"

// YAMLScan :
func YAMLScan(data, dfltRoot, pDeli string, IDs []string, dt g.DataType, OnValueFetch func(path, value, id string)) {
	switch dt {
	case g.XML:
		YAMLScanFromXMLBat(data, pDeli, IDs, OnValueFetch)
	case g.JSON:
		YAMLScanFromJSONBat(data, dfltRoot, pDeli, IDs, OnValueFetch)
	}
}

// YAMLScanFromXMLBat :
func YAMLScanFromXMLBat(xml, pDeli string, IDs []string, OnValueFetch func(path, value, id string)) {

	n, prevEnd := XMLSegsCount(xml), 0
	for i := 1; i <= n; i++ {
		nextStart := IF(i == 1, 0, prevEnd+1).(int)
		_, thisxml, _, end := XMLSegPos(S(xml).S(nextStart, ALL).V(), 1, 1)
		prevEnd = end + nextStart

		// fPf("%d SIF *****************************************************\n", i)

		_, _, idtags, _, _ := XMLScanObjects(thisxml)
		idmark := idtags[0]

		yaml := Xstr2Y(thisxml)
		yaml = YAMLJoinSplittedLines(yaml)
		info := YAMLInfo(yaml, idmark, pDeli, true)
		for _, item := range *info {
			ID := item.ID
			if IDs != nil && len(IDs) > 0 {
				ID = IDs[i-1]
			}
			OnValueFetch(item.Path, item.Value, ID)
		}
	}
}

// YAMLScanFromJSONBat :
func YAMLScanFromJSONBat(json, dfltRoot, pDeli string, IDs []string, OnValueFetch func(path, value, id string)) {

	if ok, _ := IsJSONSingle(json); ok {

		IDTag, _, _, _, _, _ := JSONObjInfo(json, dfltRoot, pDeli)
		yaml := Jstr2Y(json)
		yaml = YAMLJoinSplittedLines(yaml)
		info := YAMLInfo(yaml, IDTag, pDeli, true)
		for _, item := range *info {
			ID := item.ID
			if IDs != nil && len(IDs) > 0 {
				ID = IDs[0]
			}
			OnValueFetch(item.Path, item.Value, ID)
		}
		return
	}

	if ok, jsonType, n, eles := IsJSONArray(json); ok {

		if jsonType == JT_OBJ {
			for i := 1; i <= n; i++ {
				thisjson := eles[i-1]

				IDTag, _, _, _, _, extjson := JSONObjInfo(thisjson, dfltRoot, pDeli)
				// fPf("%d json *****************************************************\n", i)
				yaml := Jstr2Y(extjson)
				yaml = YAMLJoinSplittedLines(yaml)
				info := YAMLInfo(yaml, IDTag, pDeli, true)
				for _, item := range *info {
					ID := item.ID
					if IDs != nil && len(IDs) > 0 {
						ID = IDs[i-1]
					}
					OnValueFetch(item.Path, item.Value, ID)
				}
			}
		}

	} else {

		IDTag, _, _, _, _, extjson := JSONObjInfo(json, dfltRoot, pDeli)
		yaml := Jstr2Y(extjson)
		yaml = YAMLJoinSplittedLines(yaml)
		// ioutil.WriteFile("tempyaml.yaml", []byte(yaml), 0666 )

		info := YAMLInfo(yaml, IDTag, pDeli, true)
		for _, item := range *info {
			ID := item.ID
			if IDs != nil && len(IDs) > 0 {
				ID = IDs[0]
			}
			OnValueFetch(item.Path, item.Value, ID)
		}
	}
}

package xjy

// YAMLScan :
func YAMLScan(data, idmark string, IDs []string, dt DataType, OnValueFetch func(path, value, id string)) {
	switch dt {
	case DT_XML:
		YAMLScanFromXMLBat(data, idmark, IDs, OnValueFetch)
	case DT_JSON:
		YAMLScanFromJSONBat(data, idmark, IDs, OnValueFetch)
	}
}

// YAMLScanFromXMLBat :
func YAMLScanFromXMLBat(xml, idmark string, IDs []string, OnValueFetch func(path, value, id string)) {
	n, prevEnd := XMLSegsCount(xml), 0
	for i := 1; i <= n; i++ {
		nextStart := IF(i == 1, 0, prevEnd+1).(int)
		_, thisxml, _, end := XMLSegPos(Str(xml).S(nextStart, ALL).V(), 1, 1)
		prevEnd = end + nextStart

		// if i == n {
		// 	fPln(tag)
		// 	fPln(thisxml)
		// }
		fPf("%d SIF *****************************************************\n", i)

		yamlstr := Xstr2Y(thisxml)
		yamlstr = YAMLJoinSplittedLines(yamlstr)
		info := YAMLInfo(yamlstr, idmark, PATH_DEL, true)
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
func YAMLScanFromJSONBat(json, idmark string, IDs []string, OnValueFetch func(path, value, id string)) {
	if ok, _ := IsJSONSingle(json); ok {
		yamlstr := Jstr2Y(json)
		yamlstr = YAMLJoinSplittedLines(yamlstr)
		info := YAMLInfo(yamlstr, idmark, PATH_DEL, true)
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
				_, _, extjson := JSONWrapRoot(thisjson, "xapi")

				// if i == n {
				// 	fPln(thisjson)
				// }
				fPf("%d xapi *****************************************************\n", i)

				yamlstr := Jstr2Y(extjson)
				yamlstr = YAMLJoinSplittedLines(yamlstr)
				info := YAMLInfo(yamlstr, idmark, PATH_DEL, true)
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
		_, _, extjson := JSONWrapRoot(json, "xapi")
		yamlstr := Jstr2Y(extjson)
		yamlstr = YAMLJoinSplittedLines(yamlstr)
		
		// ioutil.WriteFile("tempyaml.yaml", []byte(yamlstr), 0666 )

		info := YAMLInfo(yamlstr, idmark, PATH_DEL, true)
		for _, item := range *info {
			ID := item.ID
			if IDs != nil && len(IDs) > 0 {
				ID = IDs[0]
			}
			OnValueFetch(item.Path, item.Value, ID)
		}
	}
}

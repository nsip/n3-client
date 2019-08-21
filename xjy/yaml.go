package xjy

import (
	"errors"

	g "github.com/nsip/n3-client/global"
)

// YAMLScan :
func YAMLScan(data, dfltRoot, pDeli string, IDs []string, dt g.DataType, OnValueFetch func(path, value, id string) error) error {
	switch dt {
	case g.XML:
		return YAMLScanFromXMLBat(data, pDeli, IDs, OnValueFetch)
	case g.JSON:
		return YAMLScanFromJSONBat(data, dfltRoot, pDeli, IDs, OnValueFetch)
	}
	return errors.New("Data Type Input Error")
}

// YAMLScanFromXMLBat :
func YAMLScanFromXMLBat(xml, pDeli string, IDs []string, OnValueFetch func(path, value, id string) error) error {

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
			if e := OnValueFetch(item.Path, item.Value, ID); e != nil {
				return e
			}
		}
	}
	return nil
}

// YAMLScanFromJSONBat :
func YAMLScanFromJSONBat(json, dfltRoot, pDeli string, IDs []string, OnValueFetch func(path, value, id string) error) error {

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
			if e := OnValueFetch(item.Path, item.Value, ID); e != nil {
				return e
			}
		}
		return nil
	}
	
	if ok, jsonType, n, eles := IsJSONArrOnFmtL0(json); ok {

		if jsonType == J_OBJ {
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
					if e := OnValueFetch(item.Path, item.Value, ID); e != nil {
						return e
					}
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
			if e := OnValueFetch(item.Path, item.Value, ID); e != nil {
				return e
			}
		}
	}

	return nil
}

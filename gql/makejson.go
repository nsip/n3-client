package gql

import (
	u "github.com/cdutwhu/go-util"
)

// JSONMake :
func JSONMake(json, path, pathDel, childDel string) string {
	if json == "" {
		json, _ = u.Str("").JSONBuild("", "", 1, path, "{}")
	}

	for _, f := range sSpl(mapStruct[sRepAll(path, "[]", "")], childDel) {

		if f == "" {
			return json
		}

		xpath := path + pathDel + f
		// fPln(xpath)
		// if xpath == "StaffPersonal ~ PersonInfo ~ Demographics ~ CountriesOfCitizenship ~ []CountryOfCitizenship" {
		// 	fPln("stop")
		// }
		// ioutil.WriteFile("debug.json", []byte(json), 0666)

		tp, tf := sRepAll(path, "[]", ""), sRepAll(f, "[]", "")

		if ok, valvers := isLeafValue(xpath); ok { //                           	              *** VALUE ***

			if okArr, _, plain := isArray(xpath); okArr && plain { //                             *** PLAIN ARRAY VALUES ***
				for _, vv := range valvers {
					json, _ = u.Str(json).JSONBuild(tp, pathDel, 1, tf, vv.value)
				}
			} else {
				for i, vv := range valvers { //                                                   ** if len(valvers) > 1, Array Object Items
					json, _ = u.Str(json).JSONBuild(tp, pathDel, len(valvers)-i, tf, vv.value) // ** if change array order, p3 -> (i+1)
				}
			}

		} else { //                                     							              *** ARRAY OR OBJECT ***

			if isObject(xpath) { //                     							              *** OBJECT ***
				if ok, n, plain := isArray(path); ok && !plain {
					for i := 1; i <= n; i++ {
						json, _ = u.Str(json).JSONBuild(tp, pathDel, i, tf, "{}")
					}
				} else {
					json, _ = u.Str(json).JSONBuild(tp, pathDel, 1, tf, "{}")
				}
			}

			if ok, n, plain := isArray(xpath); ok && !plain { //  						          *** OBJECT ARRAY FRAME ***
				for j := 0; j < n; j++ {
					json, _ = u.Str(json).JSONBuild(tp, pathDel, 1, tf, "{}")
				}
			}
		}

		// ioutil.WriteFile("debug.json", []byte(json), 0666)
		json = JSONMake(json, xpath, pathDel, childDel)
	}

	return json
}

package gql

import u "github.com/cdutwhu/go-util"

func JSONMake(json, path string) string {
	if json == "" {
		json, _ = u.Str("").JSONBuild("", "", 1, path, "{}")
	}

	for _, f := range sSpl(mapStruct[path], " + ") {

		if f == "" {
			return json
		}

		arrFlag := false
		if sHP(f, "[]") {
			f, arrFlag = f[2:], true
		}

		xpath := path + "." + f

		if valvers, ok := isEndValue(xpath); ok { //                            *** VALUE ***

			if !arrFlag { //										            *** Normal Values ***

				if len(valvers) > 1 { //                                        *** if > 1, Array Object Items ***
					// fPln(path)
					for i, vv := range valvers {
						json, _ = u.Str(json).JSONBuild(path, ".", i+1, f, vv.value)
					}
				} else {
					json, _ = u.Str(json).JSONBuild(path, ".", 1, f, valvers[0].value)
				}

			} else { //												            *** Plain Array Values ***

				content := "[]"
				if len(valvers) > 0 {
					for _, vv := range valvers {
						content += (u.Str(vv.value).MkQuotes(u.QDouble) + ",")
					}
					content = "[" + content[:len(content)-1] + "]"
				}
				json, _ = u.Str(json).JSONBuild(path, ".", 1, f, content)
			}

		} else if _, _, ok := isArr(xpath); ok { //                             *** Array, Need further process ***

			// fPln(xpath)
			json, _ = u.Str(json).JSONBuild(path, ".", 1, f, "{}")

		} else { //                                                             *** OBJECT ***

			if arrFlag { //                                                     *** Array Objects ***

				// fPln(xpath)
				mapArrObjLen[xpath] = cntChildren(xpath)

				content := "[]"
				if arrCnt, ok := isArrPath(xpath); ok {
					for i := 0; i < arrCnt; i++ {
						content += "{},"
					}
					if arrCnt > 0 {
						content = "[" + content[:len(content)-1] + "]"
					}
				}
				json, _ = u.Str(json).JSONBuild(path, ".", 1, f, content)

			} else { //                                                         *** Normal Object ***

				if repeat, ok := isArrPath(xpath); ok { //                      *** Objects in Array ***
					for i := 0; i < repeat; i++ {
						json, _ = u.Str(json).JSONBuild(path, ".", i+1, f, "{}")
					}
				} else { //                                                     *** Single Object ***
					json, _ = u.Str(json).JSONBuild(path, ".", 1, f, "{}")
				}

			}
		}

		json = JSONMake(json, xpath)
	}

	return json
}

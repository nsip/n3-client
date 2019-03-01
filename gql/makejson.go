package gql

import u "github.com/cdutwhu/go-util"

func JSONMake(json, path, pathDel, childDel string) string {
	if json == "" {
		json, _ = u.Str("").JSONBuild("", "", 1, path, "{}")
	}

	for _, f := range sSpl(mapStruct[path], childDel) {

		if f == "" {
			return json
		}

		arrFlag := false
		if sHP(f, "[]") {
			f, arrFlag = f[2:], true
		}

		xpath := path + pathDel + f
		fPln(xpath)

		if xpath == "StaffPersonal ~ PersonInfo ~ OtherNames ~ Name" {
			fPln("stop")
		}

		if valvers, ok := isEndValue(xpath); ok { //                            *** VALUE ***

			if !arrFlag { //										            *** Normal Values ***

				//   *** len(valvers) > 1, Array Object Items ***
				for i, vv := range valvers {
					json, _ = u.Str(json).JSONBuild(path, pathDel, i+1, f, vv.value)
				}

			} else { //												            *** Plain Array Values ***

				content := ""
				if len(valvers) > 0 {
					for _, vv := range valvers {
						content += (u.Str(vv.value).MkQuotes(u.QDouble) + ",")
					}
					content = "[" + content[:len(content)-1] + "]"
				} else {
					content = "[]"
				}
				json, _ = u.Str(json).JSONBuild(path, pathDel, 1, f, content)
			}

		} else if _, _, ok := isArr(xpath); ok { //                             *** Array, Need further process ***

			// fPln(xpath)
			json, _ = u.Str(json).JSONBuild(path, pathDel, 1, f, "{}")

		} else { //                                                             *** OBJECT ***

			if arrFlag { //                                                     *** Array Objects ***

				// fPln(xpath)

				content := ""
				if arrCnt, ok := isArrPath(xpath, pathDel); ok {
					if arrCnt > 0 {
						for i := 0; i < arrCnt; i++ {
							content += "{},"
						}
						content = "[" + content[:len(content)-1] + "]"
					} else {
						content = "[]"
					}
				}
				json, _ = u.Str(json).JSONBuild(path, pathDel, 1, f, content)

			} else { //                                                         *** Normal Object ***

				if repeat, ok := isArrPath(xpath, pathDel); ok { //                 *** Objects in Array ***
					for i := 0; i < repeat; i++ {
						json, _ = u.Str(json).JSONBuild(path, pathDel, i+1, f, "{}")
					}
				} else { //                                                     *** Single Object ***
					json, _ = u.Str(json).JSONBuild(path, pathDel, 1, f, "{}")
				}

			}
		}

		json = JSONMake(json, xpath, pathDel, childDel)
	}

	return json
}

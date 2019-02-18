package gql

import (
	"io/ioutil"
	"testing"

	u "github.com/cdutwhu/go-util"
)

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
			f = f[2:]
			arrFlag = true
		}

		xpath := path + "." + f

		if values, ok := isEndValue(xpath); ok { //                 *** VALUE ***

			if !arrFlag { //										*** Normal Values ***
				for i, value := range values {
					json, _ = u.Str(json).JSONBuild(path, ".", i+1, f, value)
				}
			} else { //												*** Plain Array Values ***

				content := "["
				for _, value := range values {
					content += (u.Str(value).MkQuotes(u.QDouble) + ",")
				}
				content = content[:len(content)-1] + "]"
				json, _ = u.Str(json).JSONBuild(path, ".", 1, f, content)
			}

		} else if _, _, ok := isArr(xpath); ok { //                 *** Array, Need further process ***

			json, _ = u.Str(json).JSONBuild(path, ".", 1, f, "{}")

		} else { //                                                 *** OBJECT ***

			if arrFlag { //                                         *** Array Structure ***

				content := "["
				if arrCnt, ok := isArrPath(xpath); ok {
					for i := 0; i < arrCnt; i++ {
						content += "{},"
					}
					content = content[:len(content)-1] + "]"
				}
				json, _ = u.Str(json).JSONBuild(path, ".", 1, f, content)

			} else { //                                             *** Normal Object ***

				if repeat, ok := isArrPath(xpath); ok { //          *** Objects in Array ***
					for i := 0; i < repeat; i++ {
						json, _ = u.Str(json).JSONBuild(path, ".", i+1, f, "{}")
					}
				} else { //                                         *** Single Object ***
					json, _ = u.Str(json).JSONBuild(path, ".", 1, f, "{}")
				}

			}
		}

		json = JSONMake(json, xpath)
	}

	return json
}

func TestQueryObject(t *testing.T) {
	queryObject("D3E34F41-9D75-101A-8C3D-00AA001A1656") // *** get root, mapStruct, mapValue
	json := JSONMake("", root)
	ioutil.WriteFile("./test.json", []byte(json), 0666)

	// for k, v := range mapStruct {
	// 	fPf("%-100s%s\n", k, v)
	// }
}

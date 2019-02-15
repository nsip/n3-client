package gql

import u "github.com/cdutwhu/go-util"

func jsonmaking(jsonStr, fullpath string) string {

	if stru, ok := mapStruct[fullpath]; ok {

		// marks := sSpl(fullpath, ".")
		// mark := marks[len(marks)-1]

		for _, f := range sSpl(stru, " + ") {
			xpath := fullpath + "." + f

			if isEndValue(xpath) { //                                    *** end value

				// jsonStr, _ = jumpinto(jsonStr, mark, f, mapValue[xpath])
				jsonStr, _ = u.Str(jsonStr).JSONBuild(xpath, ".", 1, f, mapValue[xpath])

			} else if _, ok := isArrStruct(xpath); ok { //         *** array

				// jsonStr, _ = jumpinto(jsonStr, mark, f, "{}")
				// arrstr := "["
				// for i := 0; i < mapArray[xpath]; i++ {
				// 	arrstr += "{},"
				// }
				// arrstr = arrstr[:len(arrstr)-1]
				// arrstr += "]"
				// jsonStr, _ = jumpinto(jsonStr, f, arrname, arrstr)

				// fPln(arrname)

			} else { //                                                  *** another object

				// jsonStr, _ = jumpinto(jsonStr, mark, f, "{}")

			}
		}
	}
	return jsonStr
}

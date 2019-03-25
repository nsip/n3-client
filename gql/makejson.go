package gql

import (
	u "github.com/cdutwhu/go-util"
)

// JSONMake :
func JSONMake(json, path, pathDel, childDel string) string {
	if json == "" {
		json, _ = u.Str("").JSONBuild("", "", path, "{}")
	}

	for _, f := range sSpl(mapStruct[sRepAll(path, "[]", "")], childDel) {

		if f == "" {
			return json
		}

		xpath := path + pathDel + f
		// fPln(xpath)

		// if xpath == "TeachingGroup ~ StudentList ~ []TeachingGroupStudent ~ Name ~ -Type" {
		// 	fPln("stop")
		// }
		// ioutil.WriteFile("debug.json", []byte(json), 0666)

		tp, tf := sRepAll(path, "[]", ""), sRepAll(f, "[]", "")
		nSeg := sCnt(tp, pathDel) + 1

		if ok, valvers := isLeafValue(xpath); ok { //                   *** VALUE ***

			if okArr, _, plain := isArray(xpath); okArr && plain { //   *** PLAIN ARRAY VALUES ***
				nValvers := len(valvers)
				for i := range valvers {
					vv := valvers[nValvers-1-i] //                      *** reverse order ***

					// *** will change to #a ~ #b ...
					indices := []int{}
					for i := 0; i < nSeg; i++ {
						indices = append(indices, 1)
					}
					// ***

					json, _ = u.Str(json).JSONBuild(tp, pathDel, tf, vv.value, indices...)
					// ioutil.WriteFile("debug.json", []byte(json), 0666) // *** DEBUG ***
				}
			} else {

				for i, vv := range valvers { // ** if len(valvers) > 1, Array Object Items **

					// *** will change to #a ~ #b ...
					indices := []int{}
					for i := 0; i < nSeg; i++ {
						indices = append(indices, 1)
					}
					if tp == "TeachingGroup ~ StudentList ~ TeachingGroupStudent ~ Name" {
						indices[len(indices)-2] = len(valvers) - i
					}
					indices[len(indices)-1] = len(valvers) - i // ** if revers order, len(valvers)-i <--> (i+1)
					// ***

					fPln(tp, " : ", tf, " : ", i)

					json, _ = u.Str(json).JSONBuild(tp, pathDel, tf, vv.value, indices...)
					// ioutil.WriteFile("debug.json", []byte(json), 0666) // *** DEBUG ***
				}
			}

		} else { //                                     				*** ARRAY OR OBJECT ***

			if isObject(xpath) { //                     				*** OBJECT ***
				if ok, n, plain := isArray(path); ok && !plain {
					for i := 1; i <= n; i++ {

						// *** will change to #a ~ #b ...
						indices := []int{}
						for i := 0; i < nSeg; i++ {
							indices = append(indices, 1)
						}
						indices[len(indices)-1] = i
						// ***

						json, _ = u.Str(json).JSONBuild(tp, pathDel, tf, "{}", indices...)
						// ioutil.WriteFile("debug.json", []byte(json), 0666) // *** DEBUG ***
					}
				} else {

					// *** will change to #a ~ #b ...
					indices := []int{}
					for i := 0; i < nSeg; i++ {
						indices = append(indices, 1)
					}
					// ***

					json, _ = u.Str(json).JSONBuild(tp, pathDel, tf, "{}", indices...)
					// ioutil.WriteFile("debug.json", []byte(json), 0666) // *** DEBUG ***
				}
			}

			if ok, n, plain := isArray(xpath); ok && !plain { //  		*** OBJECT ARRAY FRAME ***
				for j := 0; j < n; j++ {

					// *** will change to #a ~ #b ...
					indices := []int{}
					for i := 0; i < nSeg; i++ {
						indices = append(indices, 1)
					}
					// ***

					json, _ = u.Str(json).JSONBuild(tp, pathDel, tf, "{}", indices...)
					// ioutil.WriteFile("debug.json", []byte(json), 0666) // *** DEBUG ***
				}
			}
		}

		json = JSONMake(json, xpath, pathDel, childDel)
		// ioutil.WriteFile("debug.json", []byte(json), 0666) // *** DEBUG ***
	}

	return json
}

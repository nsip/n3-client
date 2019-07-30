package gql

import g "github.com/nsip/n3-client/global"

// JSONBuild : Need global vars
func JSONBuild(path string) {

	childList, ok := "", false
	if childList, ok = mStruct[sRepAll(path, "[]", "")]; !ok {
		return
	}

	for _, field := range sSpl(childList, g.DELIChild) {

		tfield := sRepAll(field, "[]", "")
		xpath := path + g.DELIPath + field
		xpath = sRepAll(xpath, "[]", "")

		if ok, valvers := isLeafValue(xpath); ok { //                                        *** PLAIN VALUE(S) ***

			// *** reverse the values ***
			LV := len(valvers)
			for i := LV/2 - 1; i >= 0; i-- {
				opp := LV - 1 - i
				valvers[i], valvers[opp] = valvers[opp], valvers[i]
			}
			// *** reverse the values ***

			if ixpaths := IPathListBymArr(xpath); len(ixpaths) > 0 { //                      *** PLAIN [ARRAY] VALUES ***
				for _, ixpath := range ixpaths {
					if okArr, nArr, plain := isArray(ixpath); okArr && plain {
						vidx := getVIdxForIPath(ixpaths, ixpath)
						ipath := S(ixpath).RmTailFromLast(g.DELIPath).V() //                 * already has '#' *
						for i := vidx; i < vidx+nArr; i++ {
							JSONMake(mIPathObj, ipath, tfield, valvers[i].value, true)
						}
					}
				}
			} else { //                                                                      *** PLAIN SINGLE VALUE ***
				if subIPathList := SubIPathListByPath(path); len(subIPathList) > 0 {
					for i := 0; i < len(subIPathList); i++ {
						ipath := subIPathList[i]
						JSONMake(mIPathObj, ipath, tfield, valvers[i].value, false)
					}
				} else {
					fs := sSpl(path, g.DELIPath)
					is := IArrMake("Strs", len(fs), "1")
					ipath := IArrStrJoinEx(Ss(fs), is.(Ss), "#", g.DELIPath)
					JSONMake(mIPathObj, ipath, tfield, valvers[0].value, false)
				}
			}

		} else { //                                                                          *** OBJECT ***

			xpath = S(xpath).Replace("[]", "").V()
			// fPln("<OBJECT>               --->", xpath)

			if ixpaths := IPathListBymArr(xpath); len(ixpaths) > 0 { //                      *** OBJECT [ARRAY] ***
				// fPln("<OBJECT ARRAY>:        --->", xpath)

				for _, ixpath := range ixpaths {
					if okArr, nArr, plain := isArray(ixpath); okArr && !plain {
						ipath := S(ixpath).RmTailFromLast(g.DELIPath).V() //                 * already has '#' *
						for i := 1; i <= nArr; i++ {
							sub := fSf("%s#%d", ixpath, i)
							JSONMake(mIPathObj, ipath, tfield, sub, true)
							mIPathSubIPaths[ixpath] = append(mIPathSubIPaths[ixpath], sub)
						}
					}
				}

			} else { //                                                                      *** OBJECT SINGLE ***

				// fPf("<OBJECT SINGLE> ---> %-30s%-30s%-30s%-30s\n", path, xpath, tfield, ipath)

				if subIPathList := SubIPathListByPath(path); len(subIPathList) > 0 {

					for i := 0; i < len(subIPathList); i++ {
						ipath := subIPathList[i]
						ixpath := ipath + g.DELIPath + tfield
						sub := ipath + g.DELIPath + tfield + "#1"
						JSONMake(mIPathObj, ipath, tfield, sub, false)
						mIPathSubIPaths[ixpath] = append(mIPathSubIPaths[ixpath], sub)
					}

				} else {

					fs := sSpl(path, g.DELIPath)
					is := IArrMake("Strs", len(fs), "1")
					ipath := IArrStrJoinEx(Ss(fs), is.(Ss), "#", g.DELIPath)
					sub := ipath + g.DELIPath + tfield + "#1"
					JSONMake(mIPathObj, ipath, tfield, sub, false)

				}
			}
		}

		JSONBuild(xpath)
	}
	return
}

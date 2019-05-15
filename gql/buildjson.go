package gql

// JSONBuild : Need global vars
func JSONBuild(path string) {

	childList, ok := "", false
	if childList, ok = mStruct[sRepAll(path, "[]", "")]; !ok {
		return
	}

	for _, field := range sSpl(childList, CHILD_DEL) {

		tfield := sRepAll(field, "[]", "")
		xpath := path + PATH_DEL + field
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
						// fPln("<LeafValue PLAIN ARRAY>:  --->", ixpath, "[", nArr, vidx, "]")
						ipath := Str(ixpath).RmTailFromLast(PATH_DEL).V() //      * already has '#' *
						for i := vidx; i < vidx+nArr; i++ {
							JSONMake(mIPathObj, ipath, tfield, valvers[i].value, true)
						}
					}
				}

			} else { //                                                                      *** PLAIN SINGLE VALUE ***

				if len(mIPathSubIPaths) > 0 {
					if subIPathList := SubIPathListByPath(path); len(subIPathList) > 0 {
						for i := 0; i < len(subIPathList); i++ {
							ipath := subIPathList[i]
							JSONMake(mIPathObj, ipath, tfield, valvers[i].value, false)
						}
					}
				} else {

					fs := sSpl(path, PATH_DEL)
					is := IArrMake("Strs", len(fs), "1")
					ipath := IArrStrJoinEx(Strs(fs), is.(Strs), "#", PATH_DEL)
					JSONMake(mIPathObj, ipath, tfield, valvers[0].value, false)
				}
			}

		} else { //                                                                          *** OBJECT ***

			xpath = sRepAll(xpath, "[]", "")
			// fPln("<OBJECT>               --->", xpath)

			if ixpaths := IPathListBymArr(xpath); len(ixpaths) > 0 { //                      *** OBJECT [ARRAY] ***
				// fPln("<OBJECT ARRAY>:        --->", xpath)

				for _, ixpath := range ixpaths {
					if okArr, nArr, plain := isArray(ixpath); okArr && !plain {
						// fPln("<OBJECT ARRAY>:        --->", ixpath, "[", nArr, "]")
						ipath := Str(ixpath).RmTailFromLast(PATH_DEL).V() //   * already has '#' *
						for i := 1; i <= nArr; i++ {
							sub := fSf("%s#%d", ixpath, i)
							mIPathSubIPaths[ixpath] = append(mIPathSubIPaths[ixpath], sub)
							JSONMake(mIPathObj, ipath, tfield, sub, true)
						}
					}
				}

			} else { //                                                                      *** OBJECT SINGLE ***

				// fPf("<OBJECT SINGLE> ---> %-30s%-30s%-30s%-30s\n", path, xpath, tfield, ipath)

				fs := sSpl(path, PATH_DEL)
				is := IArrMake("Strs", len(fs), "1")
				ipath := IArrStrJoinEx(Strs(fs), is.(Strs), "#", PATH_DEL)
				sub := ipath + PATH_DEL + tfield + "#1"
				JSONMake(mIPathObj, ipath, tfield, sub, false)
			}
		}

		JSONBuild(xpath)
	}
	return
}

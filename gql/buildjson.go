package gql

// JSONBuild : Need global vars
func JSONBuild(path string) {

	for _, field := range sSpl(mStruct[sRepAll(path, "[]", "")], CHILD_DEL) {
		if field == "" {
			return
		}

		tfield := sRepAll(field, "[]", "")
		xpath := path + PATH_DEL + field
		xpath = sRepAll(xpath, "[]", "")

		if ok, valvers := isLeafValue(xpath); ok { //                                        *** PLAIN VALUE(S) ***

			// fPln("<LeafValue>               --->", xpath)

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
						ipath := Str(ixpath).RmTailFromLast(PATH_DEL).V()
						for i := vidx; i < vidx+nArr; i++ {							
							JSONMake(mIPathObj, ipath, tfield, valvers[i].value, true)
						}
						fPln()
					}
				}

			} else { //                                                                      *** PLAIN SINGLE VALUE ***

				ipath := IF(!Str(path).Contains(PATH_DEL), path+"#1", path).(string)

				// if tfield == "description" {
				// 	fPln("debug")
				// }

				if len(mIPathSubIPaths) > 0 {
					if subIPathList := SubIPathListByPath(path); len(subIPathList) > 0 {
						for i := 0; i < len(subIPathList); i++ {
							JSONMake(mIPathObj, subIPathList[i], tfield, valvers[i].value, false)
						}
					}
				} else {
					JSONMake(mIPathObj, ipath, tfield, valvers[0].value, false)
				}
			}

		} else { //                                                                          *** OBJECT ***

			xpath = sRepAll(xpath, "[]", "")
			// fPln("<SubObject>               --->", xpath)
			if ixpaths := IPathListBymArr(xpath); len(ixpaths) > 0 { //                      *** OBJECT [ARRAY] ***
				// fPln("<SubObject ARRAY>:        --->", xpath)
				for _, ixpath := range ixpaths {
					if okArr, nArr, plain := isArray(ixpath); okArr && !plain {
						// fPln("<SubObject ARRAY>:        --->", ixpath, "[", nArr, "]")
						ipath := Str(ixpath).RmTailFromLast(PATH_DEL).V()
						for i := 1; i <= nArr; i++ {
							sub := fSf("%s#%d", ixpath, i)
							mIPathSubIPaths[ixpath] = append(mIPathSubIPaths[ixpath], sub)
							JSONMake(mIPathObj, ipath, tfield, sub, true)
						}
					}
				}
			} else { //                                                                      *** OBJECT SINGLE ***

				fPln("<SubObject SINGLE> --->", xpath)
				panic("<SubObject SINGLE> ---> unimplemented")

			}

		}

		fPln()
		JSONBuild(xpath)
	}
	return
}

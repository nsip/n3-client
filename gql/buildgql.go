package gql

// SchemaBuild : init path is root
func SchemaBuild(gql, path string) string {

	tps := sSpl(path, PATH_DEL)
	tp := tps[len(tps)-1]

	for _, f := range sSpl(mStruct[path], CHILD_DEL) {
		if f == "" {
			return gql
		}

		arrFlag := false
		if Str(f).HP("[]") {
			f, arrFlag = f[2:], true
		}

		xpath := path + PATH_DEL + f
		if ok, _ := isLeafValue(xpath); ok {
			if arrFlag {
				gql = SchemaMake(Str(gql), tp, f, "[String]")
			} else {
				gql = SchemaMake(Str(gql), tp, f, "String")
			}
		} else {
			if arrFlag {
				gql = SchemaMake(Str(gql), tp, f, "["+f+"]")
			} else {
				gql = SchemaMake(Str(gql), tp, f, f)
			}
		}

		gql = SchemaBuild(gql, xpath)
	}

	return gql
}

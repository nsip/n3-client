package gql

import g "github.com/nsip/n3-client/global"

// SchemaBuild : init path is root
func SchemaBuild(gql, path string) string {

	tps := sSpl(path, g.DELIPath)
	tp := tps[len(tps)-1]

	for _, f := range sSpl(mStruct[path], g.DELIChild) {
		if f == "" {
			return gql
		}

		arrFlag := false
		if S(f).HP("[]") {
			f, arrFlag = f[2:], true
		}

		xpath := path + g.DELIPath + f
		if ok, _ := isLeafValue(xpath); ok {
			if arrFlag {
				gql = SchemaMake(S(gql), tp, f, "[String]")
			} else {
				gql = SchemaMake(S(gql), tp, f, "String")
			}
		} else {
			if arrFlag {
				gql = SchemaMake(S(gql), tp, f, "["+f+"]")
			} else {
				gql = SchemaMake(S(gql), tp, f, f)
			}
		}

		gql = SchemaBuild(gql, xpath)
	}

	return gql
}

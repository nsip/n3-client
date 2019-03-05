package gql

import (
	u "github.com/cdutwhu/go-util"
)

// SchemaMake : init path is root
func SchemaMake(gql, path, pathDel, childDel string) string {

	tps := sSpl(path, pathDel)
	tp := tps[len(tps)-1]

	for _, f := range sSpl(mapStruct[path], childDel) {
		if f == "" {
			return gql
		}

		arrFlag := false
		if sHP(f, "[]") {
			f, arrFlag = f[2:], true
		}

		xpath := path + pathDel + f
		if ok, _ := isLeafValue(xpath); ok {
			if arrFlag {
				gql = u.Str(gql).GQLBuild(tp, f, "[String]")
			} else {
				gql = u.Str(gql).GQLBuild(tp, f, "String")
			}
		} else {
			if arrFlag {
				gql = u.Str(gql).GQLBuild(tp, f, "["+f+"]")
			} else {
				gql = u.Str(gql).GQLBuild(tp, f, f)
			}
		}

		gql = SchemaMake(gql, xpath, pathDel, childDel)
	}

	return gql
}

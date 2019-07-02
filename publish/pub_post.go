package publish

import (
	"io/ioutil"
	"os"

	g "../global"
	pp "../preprocess"
)

// mkSchemaQueryHead :
func mkSchemaQueryHead(qSchemaDir string, objects ...string) {
	for _, obj := range objects {
		fpath := qSchemaDir + obj + ".gql"
		if _, e := os.Stat(fpath); e == nil {
			return
		} else if os.IsNotExist(e) {
			schema := fSf("type QueryRoot {\n\t%s: %s\n}", obj, obj)
			ioutil.WriteFile(fpath, []byte(schema), 0666)
		} else {
			panic("mkSchemaQueryHead error")
		}
	}
}

func postpJSON(json string, IDs, Objs []string) {
	// *** save original object JSON, only for 1 object file *** //
	if len(IDs) == 1 {
		ID, root := IDs[0], Objs[0]
		_, _, json = JSONWrapRoot(json, root)
		json = pp.FmtJSONStr(json, "../preprocess/util/", "./preprocess/util/", "./")
		ioutil.WriteFile(fSf("../build/debug_pub/%s.json", ID), []byte(json), 0666)
	}

	g.RmIDsInLRU(IDs...) // *** remove id from lru cache ***
	g.RmQryIDsCache(IDs...)
	mkSchemaQueryHead(CFG.Query.SchemaDir, Objs...) // *** create gql schema query header ***
}

func postpXML(xml string, IDs, Objs []string) {
	g.RmIDsInLRU(IDs...) // *** remove id from lru cache ***
	g.RmQryIDsCache(IDs...)
	mkSchemaQueryHead(CFG.Query.SchemaDir, Objs...) // *** create gql schema query header ***
}

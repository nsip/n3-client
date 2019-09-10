package publish

import (
	"io/ioutil"
	"os"

	g "github.com/nsip/n3-client/global"
)

// mkSchemaQueryHead :
func mkSchemaQueryHead(qSchemaDir string, objects ...string) {

	// fPln("all objects count ", len(objects))
	objects = IArrRmRep(Ss(objects)).([]string)
	// fPln("type objects count ", len(objects))

	for _, obj := range objects {
		fpath := qSchemaDir + obj + ".gql"
		if _, e := os.Stat(fpath); e == nil {
			continue
		} else if os.IsNotExist(e) {
			schema := fSf("type Query {\n\t%s: %s\n}\n%s\ntype Query {\n\t%s: [%s]\n}", obj, obj, DELISchema, obj, obj)
			ioutil.WriteFile(fpath, []byte(schema), 0666)
		} else {
			panic("mkSchemaQueryHead error")
		}
	}
}

func postpJSON(ctx, json string, IDs, Objs []string) {

	// *** save original object JSON, only for 1 object file *** //
	// if len(IDs) == 1 {
	// 	ID, root := IDs[0], Objs[0]
	// 	_, _, json = JSONWrapRoot(json, root)
	// 	json = pp.FmtJSONStr(json, "../preprocess/util/", "./preprocess/util/", "./")
	// 	ioutil.WriteFile(fSf("../build/debug_pub/%s.json", ID), []byte(json), 0666)
	// }

	mkSchemaQueryHead(g.Cfg.Query.SchemaDir, Objs...) // *** create gql schema query header ***
}

func postpXML(ctx, xml string, IDs, Objs []string) {
	mkSchemaQueryHead(g.Cfg.Query.SchemaDir, Objs...) // *** create gql schema query header ***
}

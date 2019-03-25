package gql

import (
	"io/ioutil"
	"testing"
)

func TestQueryObject(t *testing.T) {

	objID := "E6271E2D-700E-4286-AA81-58FEF49060C1"
	queryObject(objID, "sif") //                                 *** get root, mapStruct, mapValue ***

	fPln(root)
	fPln("<------------------------------>")

	json := JSONMake("", root, pathDel, childDel)
	json = sRepAll(json, `"-`, `"`)
	json = sRepAll(json, `"#`, `"`)
	ioutil.WriteFile(fSf("./yield/%s.json", objID), []byte(json), 0666)

	schema := SchemaMake("", root, pathDel, childDel)
	schema = sRepAll(schema, "\t-", "\t")
	schema = sRepAll(schema, "\t#", "\t")
	ioutil.WriteFile(fSf("./yield/%s.gql", objID), []byte(schema), 0666)

	// for k, v := range mapStruct {
	// 	fPf("%-100s%s\n", k, v)
	// }
}

package gql

import (
	"io/ioutil"
	"testing"
)

func TestQueryObject(t *testing.T) {

	// D3E34F41-9D75-101A-8C3D-00AA001A1656
	// D7BF9628-14ED-491E-AC96-93C54DA35FD1
	// D3E34F41-9D75-101A-8C3D-00AA001A1652
	// D3E34F41-8965-101A-8C3D-00AA001A1656
	queryObject("D3E34F41-9D75-101A-8C3D-00AA001A1652", "xapi") // *** get root, mapStruct, mapValue

	fPln(root)
	fPln("<------------------------------>")

	json := JSONMake("", root, pathDel, childDel)
	json = sRepAll(json, `"-`, `"`)
	json = sRepAll(json, `"#`, `"`)
	ioutil.WriteFile("./yield/xapi.json", []byte(json), 0666)

	schema := SchemaMake("", root, pathDel, childDel)
	schema = sRepAll(schema, "\t-", "\t")
	schema = sRepAll(schema, "\t#", "\t")
	ioutil.WriteFile("./yield/schema_xapi.gql", []byte(schema), 0666)

	// for k, v := range mapStruct {
	// 	fPf("%-100s%s\n", k, v)
	// }
}

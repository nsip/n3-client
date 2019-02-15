package gql

import (
	"testing"

	u "github.com/cdutwhu/go-util"
)

func TestQueryObject(t *testing.T) {
	queryObject("D3E34F41-9D75-101A-8C3D-00AA001A1656") // *** get root, mapStruct, mapValue
	fPln(root)

	// for k, v := range mapStruct {
	// 	fPf("%-100s%s\n", k, v)
	// }

	json, ok := u.Str("").JSONBuild("", "", 1, root, "{}")
	fPln(json, ok)

	path := root
	for _, f := range sSpl(mapStruct[path], " + ") {

		path = path + "." + f
		json = jsonmaking(json, path)

		// 	// for _, ff := range sSpl(mapStruct[path], " + ") {
		// 	// 	path = path + "." + ff
		// 	// 	jsonStr = jsonmaking(jsonStr, path)

		// 	// 	for _, fff := range sSpl(mapStruct[path], " + ") {
		// 	// 		path = path + "." + fff
		// 	// 		jsonStr = jsonmaking(jsonStr, path)
		// 	// 	}

		// 	// }
	}

	fPln(json)
}

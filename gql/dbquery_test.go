package gql

import (
	"io/ioutil"
	"testing"
)

func TestQueryObject(t *testing.T) {

	// D3E34F41-9D75-101A-8C3D-00AA001A1656
	// D7BF9628-14ED-491E-AC96-93C54DA35FD1

	queryObject("D3E34F41-9D75-101A-8C3D-00AA001A1656") // *** get root, mapStruct, mapValue
	json := JSONMake("", root)
	ioutil.WriteFile("./test.json", []byte(json), 0666)

	// for k, v := range mapStruct {
	// 	fPf("%-100s%s\n", k, v)
	// }
}

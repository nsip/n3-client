package gql

import (
	"io/ioutil"
	"testing"
)

// var ID = "738F4DF5-949F-4380-8186-8252440A6F6F"
var ID = "1a723f08-5cee-4ad7-8a48-68e7bda480fd"

func TestBuildInfoFromID(t *testing.T) {
	if data := GetInfoFromID("JSON", ID); data != "" {
		ioutil.WriteFile(fSf("./debug/%s.json", ID), []byte(data), 0666)
	}
	if data := GetInfoFromID("SCHEMA", ID); data != "" {
		ioutil.WriteFile(fSf("./debug/%s.gql", ID), []byte(data), 0666)
	}
}

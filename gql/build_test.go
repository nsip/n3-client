package gql

import (
	"io/ioutil"
	"testing"
)

// var ID = "738F4DF5-949F-4380-8186-8252440A6F6F"
var ID = "89d00298-b56d-4567-8d63-7fd02a2ff528"

func TestBuildInfoFromID(t *testing.T) {
	if data := GetInfoFromID("JSON", ID); data != "" {
		ioutil.WriteFile(fSf("./debug/%s.json", ID), []byte(data), 0666)
	}
	if data := GetInfoFromID("SCHEMA", ID); data != "" {
		ioutil.WriteFile(fSf("./debug/%s.gql", ID), []byte(data), 0666)
	}
}

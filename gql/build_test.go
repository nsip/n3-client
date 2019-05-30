package gql

import (
	"io/ioutil"
	"testing"
)

// var ID = "738F4DF5-949F-4380-8186-8252440A6F6F"
var ID = "BF89BE42-ADC9-4F43-9740-A45A737BDDF3"

func TestBuildInfoFromID(t *testing.T) {
	if data := GetInfoFromID("JSON", ID); data != "" {
		ioutil.WriteFile(fSf("./debug/%s.json", ID), []byte(data), 0666)
	}
	if data := GetInfoFromID("SCHEMA", ID); data != "" {
		ioutil.WriteFile(fSf("./debug/%s.gql", ID), []byte(data), 0666)
	}
}

package gql

import (
	"io/ioutil"
	"testing"
)

var ID = "1b0cea1c-1b25-4765-9666-b6414bc07c34"

func TestBuildInfoFromID(t *testing.T) {
	if data := GetInfoFromID("JSON", ID); data != "" {
		ioutil.WriteFile(fSf("./debug/%s.json", ID), []byte(data), 0666)
	}
	if data := GetInfoFromID("SCHEMA", ID); data != "" {
		ioutil.WriteFile(fSf("./debug/%s.gql", ID), []byte(data), 0666)
	}
	if data := GetInfoFromID("QRYTXT", ID); data != "" {
		ioutil.WriteFile(fSf("./debug/%s.txt", ID), []byte(data), 0666)
	}
}

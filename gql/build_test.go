package gql

import (
	"io/ioutil"
	"testing"
)

// var ID = "2c16f595-c0e7-4779-bed1-fb6df9e495b5"
var ID = "4947ED1F-1E94-4850-8B8F-35C653F51E9C"

func TestBuildJSON(t *testing.T) {
	json := GetJSONFromID(ID)
	ioutil.WriteFile(fSf("./debug/%s.json", ID), []byte(json), 0666)
}

func TestBuildSchema(t *testing.T) {
	schema := GetSchemaFromID(ID)
	ioutil.WriteFile(fSf("./debug/%s.gql", ID), []byte(schema), 0666)
}

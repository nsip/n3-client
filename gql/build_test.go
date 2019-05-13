package gql

import (
	"io/ioutil"
	"testing"
)

func TestBuildJSON(t *testing.T) {
	json := GetJSONFromID("c58f2f19-32ed-4258-a71b-6506b2a2f33b")
	ioutil.WriteFile("debugjson.gql", []byte(json), 0666)
}

func TestBuildSchema(t *testing.T) {
	schema := GetSchemaFromID("c58f2f19-32ed-4258-a71b-6506b2a2f33b")
	ioutil.WriteFile("debugschema.gql", []byte(schema), 0666)
}

package gql

import (
	"io/ioutil"
	"testing"
)

func TestGQL(t *testing.T) {
	objID := "c58f2f19-32ed-4258-a71b-6506b2a2f33b"
	queryBytes, _ := ioutil.ReadFile("./query.txt")
	result := GQuery([]string{objID}, fSf("type QueryRoot {\n\txapi: %s\n}", "xapi"), string(queryBytes), map[string]interface{}{}, []string{})
	ioutil.WriteFile("./yield/result.json", []byte(result), 0666)
}

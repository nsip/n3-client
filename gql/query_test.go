package gql

import (
	"io/ioutil"
	"testing"
)

func TestGQL(t *testing.T) {
	objID := "79461ce8-8fc7-4644-8aa7-e9cf82373bae"
	queryBytes, _ := ioutil.ReadFile("./query.txt")
	querySchema := fSf("type QueryRoot {\n\txapi: %s\n}", "xapi") //           *** content should be related to resolver path ***
	result := Query([]string{objID}, querySchema, string(queryBytes), map[string]interface{}{}, []string{})
	ioutil.WriteFile("./yield/result.json", []byte(result), 0666)
}

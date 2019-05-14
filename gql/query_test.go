package gql

import (
	"io/ioutil"
	"testing"
)

func TestGQL(t *testing.T) {
	objID := "1f5076bf-2a64-40b6-8fbc-838ccf941b2e"
	queryBytes, _ := ioutil.ReadFile("./query.txt")
	querySchema := fSf("type QueryRoot {\n\txapi: %s\n}", "xapi") //           *** content should be related to resolver path ***
	result := Query([]string{objID}, querySchema, string(queryBytes), map[string]interface{}{}, []string{})
	ioutil.WriteFile("./yield/result.json", []byte(result), 0666)
}

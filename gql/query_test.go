package gql

import (
	"io/ioutil"
	"testing"
)

func TestGQL(t *testing.T) {
	objID := "1a723f08-5cee-4ad7-8a48-68e7bda480fd"
	queryBytes, _ := ioutil.ReadFile("./queryC.txt") //               *** change ***
	querySchema := fSf("type QueryRoot {\n\txapi: %s\n}", "xapi") //  *** content should be related to resolver path ***
	result := Query([]string{objID}, querySchema, string(queryBytes), map[string]interface{}{}, []string{})
	ioutil.WriteFile("./yield/result.json", []byte(result), 0666)
}

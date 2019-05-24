package gql

import (
	"io/ioutil"
	"testing"
)

func TestGQL(t *testing.T) {
	objID := "9b8405f8-d6b0-4062-a8bc-16b3b69efdaf"
	queryBytes, _ := ioutil.ReadFile("./queryC.txt") //               *** change ***
	querySchema := fSf("type QueryRoot {\n\txapi: %s\n}", "xapi") //  *** content should be related to resolver path ***
	result := Query([]string{objID}, querySchema, string(queryBytes), map[string]interface{}{}, []string{})
	ioutil.WriteFile("./yield/result.json", []byte(result), 0666)
}

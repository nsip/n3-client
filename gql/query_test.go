package gql

import (
	"io/ioutil"
	"testing"
)

func TestGQL(t *testing.T) {
	objID := "a1bec29f-2dfc-418f-b686-61fa42f2c80a"
	queryBytes, _ := ioutil.ReadFile("./queryC.txt") //                        *** change ***
	querySchema := fSf("type QueryRoot {\n\txapi: %s\n}", "xapi") //           *** content should be related to resolver path ***
	result := Query([]string{objID}, querySchema, string(queryBytes), map[string]interface{}{}, []string{})
	ioutil.WriteFile("./yield/result.json", []byte(result), 0666)
}

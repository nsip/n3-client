package gql

import (
	"io/ioutil"
	"testing"
)

func TestGQL(t *testing.T) {
	objID := "4947ED1F-1E94-4850-8B8F-35C653F51E9C"
	queryBytes, _ := ioutil.ReadFile("./query.txt")
	querySchema := fSf("type QueryRoot {\n\txapi: %s\n}", "xapi") //           *** content should be related to resolver path ***
	result := Query([]string{objID}, querySchema, string(queryBytes), map[string]interface{}{}, []string{})
	ioutil.WriteFile("./yield/result.json", []byte(result), 0666)
}

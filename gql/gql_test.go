package gql

import (
	"io/ioutil"
	"testing"
)

func TestGQL(t *testing.T) {
	objID := "1BC3EAB7-3E48-4371-8C14-6D1E67BEBD6D" // sif
	queryBytes, _ := ioutil.ReadFile("./query.txt")
	query := string(queryBytes)
	result, _ := GQuery(objID, "sif", fSf("type QueryRoot {\n\troot: %s\n}", "TeachingGroup"), query, map[string]interface{}{}, []string{}) // TeachingGroup ???
	ioutil.WriteFile("./yield/result.json", []byte(result), 0666)
}

package gql

import (
	"io/ioutil"
	"testing"

	q "../query"
)

func TestGQL(t *testing.T) {
	objID := "ca669951-9511-4e53-ae92-50845d3bdcd6"
	_, _, o, _ := q.Data(objID, "") //                 *** get root ***
	if len(o) > 0 {
		root := o[0]
		querySchema, _ := ioutil.ReadFile(fSf("./qSchema/%s.gql", root)) // *** file content must be related to resolver path ***
		queryBytes, _ := ioutil.ReadFile(fSf("./qTxt/%s.txt", root))     // *** change ***
		result := Query([]string{objID}, string(querySchema), string(queryBytes), map[string]interface{}{}, []string{})
		ioutil.WriteFile(fSf("./yield/%s.json", objID), []byte(result), 0666)
		return
	}
	fPln("wrong objID")
}

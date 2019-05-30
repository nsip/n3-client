package gql

import (
	"io/ioutil"
	"testing"
)

func TestGQL(t *testing.T) {
	objID := "A759FF45-4ABD-4A59-B31B-BB0D3CA66ADC"
	queryBytes, _ := ioutil.ReadFile("./qTxt/SchoolCourseInfo.txt")     // *** change ***
	querySchema, _ := ioutil.ReadFile("./qSchema/SchoolCourseInfo.gql") // *** content should be related to resolver path ***
	result := Query([]string{objID}, string(querySchema), string(queryBytes), map[string]interface{}{}, []string{})
	ioutil.WriteFile(fSf("./yield/%s.json", objID), []byte(result), 0666)
}

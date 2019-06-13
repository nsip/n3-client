package preprocess

import (
	"io/ioutil"
	"testing"
)

func TestCheck(t *testing.T) {
	file := "../inbound/hsie/history/stage1/overview.json"
	json := string(Must(ioutil.ReadFile(file)).([]byte))
	fPln("hasColonInValue", hasColonInValue(json))
	fPln("hasSQuoteInValue", hasSQuoteInValue(json))
	// fPln(hasHyphen(string(bytes)))
	json1 := FmtJSONFile("../"+file, "./util/")
	json2 := FmtJSONStr(json, "./util/")
	PC(json1 != json2, fEf("[error in FmtJSONFile or FmtJSONStr]"))
	fPln(json1)
	fPln(" ***** ")
	fPln(json2)
}

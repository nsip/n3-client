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
	fPln(FmtJSONFile("../"+file, "./util/"))	
	fPln(" ***** ")
	// fPln(FmtJSONStr(json, "./util/"))
}

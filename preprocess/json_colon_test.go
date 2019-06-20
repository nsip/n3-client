package preprocess

import (
	"io/ioutil"
	"testing"
)

func TestCheck(t *testing.T) {
	file := "../inbound/hsie/history/stage4/content.json"
	json := string(Must(ioutil.ReadFile(file)).([]byte))

	// fPln(json)

	json1 := FmtJSONFile("../"+file, "./util/")
	fPln("HasColonInValue json1", HasColonInValue(json1), "\n")
	json1 = RplcColonInValue(json1, "^1m$")

	json2 := FmtJSONStr(json, "./util/")
	fPln("HasColonInValue json2", HasColonInValue(json2), "\n")
	json2 = RplcColonInValue(json2, "^1m$")

	PC(json1 != json2, fEf("[error in FmtJSONFile or FmtJSONStr]"))

	fPln(json1)
	fPln(" ***** ")
	fPln(json2)
}

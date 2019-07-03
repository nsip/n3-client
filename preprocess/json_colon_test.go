package preprocess

import (
	"io/ioutil"
	"testing"
)

func TestCheck(t *testing.T) {
	// file := "../inbound/hsie/history/stage4/content.json"
	file := "./sample.json"
	json := string(must(ioutil.ReadFile(file)).([]byte))

	ascii, _ := UTF8ToASCII(json)
	if !ascii {
		fPln("UTF8")
	}

	fPln(json)

	json1 := FmtJSONFile("../"+file, "./util/")
	fPln("HasColonInValue json1", HasColonInValue(json1), "\n")
	json11 := RplcValueColons(json1)

	json2 := FmtJSONStr(json, "./util/")
	fPln("HasColonInValue json2", HasColonInValue(json2), "\n")
	json22 := RplcValueColons(json2)

	pc(json11 != json22, fEf("1 [error in FmtJSONFile or FmtJSONStr]"))

	json111 := ASCIIToOri(json11)
	json222 := ASCIIToOri(json22)

	// fPln(json11)

	// fPln(json1)
	// fPln(" ***** ")
	// fPln(json2)
	// fPln(" ***** ")
	// fPln(json11)
	// fPln(" ***** ")
	// fPln(json22)

	pc(json222 != json111, fEf("2 [error in FmtJSONFile or FmtJSONStr]"))
	pc(json1 != json111, fEf("3 [error in FmtJSONFile or FmtJSONStr]"))
}

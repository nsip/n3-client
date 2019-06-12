package preprocess

import (
	"io/ioutil"
	"testing"
)

func TestCheck(t *testing.T) {
	bytes := Must(ioutil.ReadFile("../inbound/hsie/geography/stage2/content.json")).([]byte)
	fPln(hasColonInValue(string(bytes)))
	fPln(hasSQuoteInValue(string(bytes)))
	// fPln(hasHyphen(string(bytes)))

	data := FmtJSONFile("../xjy/files/content.json", "./")
	fPln(data)
}

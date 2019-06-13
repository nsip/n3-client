package preprocess

import (
	"testing"
)

func TestJQ(t *testing.T) {
	// data := string(Must(ioutil.ReadFile("../xjy/files/content.json")).([]byte))
	// data = FmtJSONStr(data, "./util/")
	// fPln(data)
	fPln(" **************** ")
	fPln(FmtJSONFile("../../xjy/files/content.json", "./util/"))
	fPln(FmtJSONFile("../sample.json", "./util/"))
}

package preprocess

import (
	"io/ioutil"
	"testing"
)

func TestJQ(t *testing.T) {
	data := string(Must(ioutil.ReadFile("../xjy/files/content.json")).([]byte))
	fPln(FmtJSONStr(data, "./util/"))
	fPln(" **************** ")
	fPln(FmtJSONFile("../../xjy/files/content.json", "./util/"))
	fPln(FmtJSONFile("../sample.json", "./util/"))
}

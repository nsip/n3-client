package preprocess

import (
	"io/ioutil"
	"testing"
)

func TestJQ(t *testing.T) {
	// data := string(must(ioutil.ReadFile("../xjy/files/content.json")).([]byte))
	// fPln(FmtJSONStr(data, "./util/"))
	// fPln(" **************** ")
	// fPln(FmtJSONFile("../../xjy/files/content.json", "./util/"))
	// fPln()
	jqrst := FmtJSONFile("./xapi.json", "./util/")
	ioutil.WriteFile("./util/jqrst.json", []byte(jqrst), 0666)
}

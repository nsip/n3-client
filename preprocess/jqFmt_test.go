package preprocess

import (
	"io/ioutil"
	"testing"
)

func TestJQ(t *testing.T) {

	data := string(Must(ioutil.ReadFile("../xjy/files/content.json")).([]byte))
	data = FmtJSONStr(data, "./")
	fPln(data)
	fPln(" **************** ")
	//fPln(FmtJSONFile("../xjy/files/content.json", "./"))
}

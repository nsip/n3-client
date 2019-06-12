package send

import "testing"

func TestJQ(t *testing.T) {
	fPln(fmtJSONStr(`{"name" : [8   , 0 , "abc"] }`, "../util/"))
	fPln(" **************** ")
	fPln(fmtJSONFile("../util/sample.json", "../util/"))
}

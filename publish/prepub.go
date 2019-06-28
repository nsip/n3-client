package publish

import (
	"io/ioutil"

	pp "../preprocess"
)

func prepJSON(json string) string {

	// json = pp.FmtJSONStr(json, "../preprocess/util/", "./preprocess/util/", "./") //      *** format json string ***

	ioutil.WriteFile("./debug_pub/temp.json", []byte(json), 0666)
	// json = pp.FmtJSONFile("../../publish/debug_pub/temp.json", "../preprocess/util/", "./preprocess/util/", "./") // *** unit test ***
	json = pp.FmtJSONFile("../../debug_pub/temp.json", "../preprocess/util/", "./preprocess/util/", "./") // *** exe ***
	ioutil.WriteFile("./debug_pub/tempFmt.json", []byte(json), 0666)

	if pp.HasColonInValue(json) {
		json = pp.RplcValueColons(json) //                        *** deal with <:> ***
	}
	if ascii, ajson := UTF8ToASCII(json); !ascii { //               *** convert to ASCII ***
		fPln("is utf8")
		return ajson
	}
	return json
}

func prepXML(xml string) string {
	if ascii, axml := UTF8ToASCII(xml); !ascii { //                 *** convert to ASCII ***
		fPln("is utf8")
		return axml
	}
	return xml
}
